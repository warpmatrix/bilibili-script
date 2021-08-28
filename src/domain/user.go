package domain

import (
	"main/src/client"
	log "main/src/logger"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	Name  string `json:"uname" mapstructure:"uname"`
	Mid   int    `json:"mid" mapstructure:"mid"`
	Level level  `json:"level_info" mapstructure:"level_info"`
	// the number of coins
	Money  int    `json:"money" mapstructure:"money"`
	Vip    vip    `json:"vip" mapstructure:"vip"`
	Wallet wallet `json:"wallet" maptructure:"wallet"`
}

var ptr *User

func GetUserInfo() (*User, error) {
	if ptr != nil {
		return ptr, nil
	}
	url := "https://api.bilibili.com/x/web-interface/nav"
	data, err := client.RecData(client.Get(url))
	if err != nil {
		return nil, err
	}
	ptr, err = parseUserInfo(data.(map[string]interface{}))
	return ptr, err
}

func parseUserInfo(data map[string]interface{}) (*User, error) {
	var u User
	if err := mapstructure.Decode(data, &u); err != nil {
		return nil, err
	}
	if remain := u.Vip.Remain; remain != nil {
		u.Vip.text = remain["label"].(map[string]interface{})["text"].(string)
	}
	return &u, nil
}

func (u *User) PrintInfo() {
	log.Info("【用户名】:", u.Name)
	log.Info("【硬币数量】:", u.Money)
	log.Info("【当前等级】:", u.Level.CurLevel)
	log.Info("【当前经验】:", u.Level.CurExp)
	switch t := u.Level.NextExp.(type) {
	case float64:
		log.Info("【距离下一级的经验】：", t)
	case string:
		log.Info("【距离一下级的经验】：当前已经是最高级")
	default:
		log.Info("【距离一下级的经验】：类型转换错误")
	}
	log.Info("【大会员类型】：", u.Vip.text)
	log.Info("【剩余 b 币卷数量】：", u.Wallet.BcoinCoupon)
}

type level struct {
	CurExp   int `json:"current_exp" mapstructure:"current_exp"`
	CurLevel int `json:"current_level" mapstructure:"current_level"`
	// if the level is six, then nextExp is a string ("--")
	// else nextExp is an int
	NextExp interface{} `json:"next_exp" mapstructure:"next_exp"`
}

type vip struct {
	Typ    int                    `json:"type" mapstructure:"type"`
	St     int                    `json:"status" mapstructure:"status"`
	Remain map[string]interface{} `mapstructure:",remain"`
	text   string
}

type wallet struct {
	BcoinCoupon float64 `json:"coupon_balance" mapstruture:"coupon_balance"`
}
