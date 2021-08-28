package task

import (
	"fmt"
	"main/src/client"
	"main/src/domain"
	log "main/src/logger"
	"net/url"
	"strconv"
	"time"
)

// first level configuration name: vip
func init() {
	config := domain.LoadCfg("vip").(map[string]interface{})
	if config == nil {
		log.Info("用户未配置大会员相关的配置")
		return
	}
	if cfg := config["privilege"]; cfg != nil {
		recvPriv.config = cfg
		taskList = append(taskList, recvPriv)
	}
	if cfg := config["bcoinUsage"]; cfg != nil {
		useBcoin.config = cfg
		taskList = append(taskList, useBcoin)
	}
}

var recvPriv = &task{
	name: "领取大会员权益",
	initFunc: func(t *task) error {
		privs, isArr := t.config.([]string)
		if !isArr {
			return fmt.Errorf("字段类型错误")
		}
		cfg := []int{}
		for _, priv := range privs {
			switch priv {
			case "bcoin":
				cfg = append(cfg, 1)
			case "coupon":
				cfg = append(cfg, 2)
			default:
				log.Warning("未识别的字段：", priv, t.checkMsg)
			}
		}
		t.config = cfg
		return nil
	},
	checkMsg:   "privilege 配置的字段应为 bcoin 或 coupon 的组合",
	defaultCfg: []string{"bcoin", "coupon"},
	impl: func(t *task) error {
		if user, err := domain.GetUserInfo(); err != nil {
			return err
		} else if !(user.Vip.Typ == 2 && user.Vip.St == 1) {
			return fmt.Errorf("该用户不是年度大会员，不能领取年度大会员权益")
		}
		if today, day := time.Now().Day(), 1; today != day {
			t.result = fmt.Sprintf("今天是 %d 日，将在 %d 日领取大会员权益", today, day)
			return nil
		}
		getPrivUrl := "https://api.bilibili.com/x/vip/privilege/my"
		data, err := client.RecData(client.Get(getPrivUrl))
		if err != nil {
			return err
		}
		privs := map[int]struct{}{}
		for _, i := range data.(map[string]interface{})["list"].([]interface{}) {
			if st := i.(map[string]interface{})["state"].(int); st == 0 {
				typ := i.(map[string]interface{})["type"].(int)
				privs[typ] = struct{}{}
			}
		}
		params := make(url.Values)
		params.Add("csrf", client.GetBiliJct())
		for _, typ := range t.config.([]int) {
			if _, ready := privs[typ]; !ready {
				continue
			}
			params.Set("type", strconv.Itoa(typ))
			url := "https://api.bilibili.com/x/vip/privilege/receive"
			if err := client.CheckCode(client.PostForm(url, params)); err != nil {
				return err
			}
		}
		t.result = "领取大会员权益成功"
		return nil
	},
}

var useBcoin = &task{
	name: "使用 b 币卷",
	initFunc: func(t *task) error {
		usage, isStr := t.config.(string)
		if !isStr {
			return fmt.Errorf("字段类型错误")
		} else if !(usage == "liveGoods" || usage == "charge") {
			return fmt.Errorf("无法识别的字段：%v", usage)
		}
		return nil
	},
	impl: func(t *task) error {
		if today, day := time.Now().Day(), 28; today != day {
			t.result = fmt.Sprintf("今天是 %d 日，将在 %d 日领取大会员权益", today, day)
			return nil
		}
		user, err := domain.GetUserInfo()
		if err != nil {
			return err
		}
		params := make(url.Values)
		params.Add("csrf", client.GetBiliJct())
		switch t.config.(string) {
		case "liveGoods":
			return chargeLiveGoods(t, user, params)
		case "charge":
			return chargeBattery(t, user, params)
		}
		return nil
	},
	checkMsg:   "bcoinUsage 配置的字段应为 liveGoods 或 charge 的字符串",
	defaultCfg: "liveGoods",
}

func chargeLiveGoods(t *task, user *domain.User, params url.Values) error {
	coupon := int(user.Wallet.BcoinCoupon)
	// 1 个 b 币对应 10 个直播电池 Set 1, 2021
	params.Add("pay_bp", strconv.Itoa(coupon*10))
	params.Add("context_id", "1")
	params.Add("context_type", "11")
	params.Add("goods_id", "1")
	params.Add("goods_num", strconv.Itoa(coupon))
	url := "https://api.live.bilibili.com/xlive/revenue/v1/order/createOrder"
	if err := client.CheckCode(client.PostForm(url, params)); err != nil {
		return err
	}
	t.result = fmt.Sprintf("成功兑换直播 %d 电池", coupon*10)
	return nil
}

func chargeBattery(t *task, user *domain.User, params url.Values) error {
	coupon := int(user.Wallet.BcoinCoupon)
	if coupon < 2 {
		return fmt.Errorf("b 币卷数量少于 2，低于充电电池下限 20")
	}
	params.Add("bp_num", strconv.Itoa(coupon))
	params.Add("is_bp_remains_prior", "true")
	params.Add("up_mid", strconv.Itoa(user.Mid))
	params.Add("oid", strconv.Itoa(user.Mid))
	params.Add("otype", "up")
	url := "https://api.bilibili.com/x/ugcpay/web/v2/trade/elec/pay/quick"
	data, err := client.RecData(client.PostForm(url, params))
	if err != nil {
		return err
	}
	switch data.(map[string]interface{})["status"].(int) {
	case 4:
		t.result = "用户充电成功"
	case -4:
		return fmt.Errorf("b 币数量不足，无法完成充电")
	default:
		return fmt.Errorf("未能识别错误原因，充电失败")
	}
	return nil
}
