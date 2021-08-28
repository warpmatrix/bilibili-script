package task

import (
	"fmt"
	"main/src/client"
	"main/src/domain"
	log "main/src/logger"
	"math"
	"math/rand"
	"net/url"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

// first level configuration name: video
func init() {
	config := domain.LoadCfg("video").(map[string]interface{})
	if config == nil {
		log.Info("用户未配置视频相关的配置")
		return
	}
	// second level configuration name: daily
	for _, daily := range config["daily"].([]interface{}) {
		switch daily {
		case "watch":
			taskList = append(taskList, watchVideo)
		case "share":
			taskList = append(taskList, shareVideo)
		default:
			log.Warning("daily 中应填写 watch 或 share 的组合，无法识别的字段：", daily)
		}
	}
	// second level configuration name: coin
	if cfg := config["coin"]; cfg != nil {
		coinVideo.config = cfg
		taskList = append(taskList, coinVideo)
	}
}

var watchVideo = &task{
	name: "模拟观看视频",
	impl: func(t *task) error {
		if rd, err := getReward(); err != nil {
			return fmt.Errorf("获取用户信息失败：%s", err)
		} else if rd.IsWatch {
			t.result = "该账户今天已观看视频"
			return nil
		}
		videos, err := getVideos(6)
		if err != nil {
			return fmt.Errorf("获取推荐视频失败：%s", err)
		}
		idx := rand.Intn(len(videos))
		time := rand.Intn(videos[idx].Duration-2) + 2
		params := make(url.Values)
		params.Add("aid", strconv.Itoa(videos[idx].Aid))
		params.Add("cid", strconv.Itoa(videos[idx].Cid))
		params.Add("progres", strconv.Itoa(time))
		params.Add("csrf", client.GetBiliJct())
		url := "https://api.bilibili.com/x/v2/history/report"
		if err := client.CheckCode(client.PostForm(url, params)); err != nil {
			return err
		}
		t.result = "成功模拟观看视频"
		return nil
	},
}

var shareVideo = &task{
	name: "分享视频",
	impl: func(t *task) error {
		if rd, err := getReward(); err != nil {
			return fmt.Errorf("获取用户信息失败：%s", err)
		} else if rd.IsShare {
			t.result = "该账户今天已分享视频"
			return nil
		}
		videos, err := getVideos(6)
		if err != nil {
			return fmt.Errorf("获取推荐视频失败：%s", err)
		}
		idx := rand.Intn(len(videos))
		params := make(url.Values)
		params.Add("bvid", videos[idx].Bvid)
		params.Add("csrf", client.GetBiliJct())
		url := "https://api.bilibili.com/x/web-interface/share/add"
		if err := client.CheckCode(client.PostForm(url, params)); err != nil {
			return err
		}
		t.result = "成功模拟分享视频"
		return nil
	},
}

type coinCfg struct {
	Num  int `mapstructure:"num"`
	Like bool `mapstructure:"like"`
}

var coinVideo = &task{
	name: "视频投币",
	initFunc: func(t *task) error {
		mapstructure.Decode(t.defaultCfg, &t.cfgMap)
		switch t.config.(type) {
		case int:
			if err := t.setCfg("num", t.config.(int)); err != nil {
				log.Warning(err)
			}
		case map[string]interface{}:
			for key, val := range t.config.(map[string]interface{}) {
				if err := t.setCfg(key, val); err != nil {
					log.Warning(err)
				}
			}
		default:
			return fmt.Errorf("coin 配置的字段应为整数或键值对")
		}
		var cfg coinCfg
		mapstructure.Decode(t.cfgMap, &cfg)
		t.config = cfg
		return nil
	},
	defaultCfg: coinCfg{Num: 5, Like: false},
	checkFuncs: map[string]func(interface{}) error{
		"num": func(i interface{}) error {
			num, isInt := i.(int)
			if !isInt {
				return fmt.Errorf("num 配置的字段应为整数")
			}
			if !(0 <= num && num <= 5) {
				return fmt.Errorf("num 配置的字段应为 0-5")
			}
			return nil
		},
		"like": func(i interface{}) error {
			_, isBool := i.(bool)
			if !isBool {
				return fmt.Errorf("like 配置的字段应为 true 或 false")
			}
			return nil
		},
	},
	impl: func(t *task) error {
		rd, err := getReward()
		if err != nil {
			return fmt.Errorf("获取用户信息失败：%s", err)
		}
		reqCoinNum := t.config.(coinCfg).Num - rd.Coin/10
		if reqCoinNum <= 0 {
			t.result = "该账户今天已完成投币"
			return nil
		}
		videos, err := getVideos(6)
		if err != nil {
			return fmt.Errorf("获取推荐视频失败：%s", err)
		}
		idxs := make([]int, len(videos))
		for i := range idxs {
			idxs[i] = i
		}
		user, err := domain.GetUserInfo()
		if err != nil {
			return err
		}
		for reqCoinNum > 0 && len(idxs) > 0 && user.Money > 0 {
			i := rand.Intn(len(idxs))
			idx := idxs[i]
			idxs[i] = idxs[len(idxs)-1]
			idxs = idxs[:len(idxs)-1]
			getCoinUrl := "http://api.bilibili.com/x/web-interface/archive/coins"
			params := [][]string{{"bvid", videos[idx].Bvid}}
			gotData, err := client.RecData(client.GetWithParams(getCoinUrl, params))
			if err != nil {
				return err
			}
			if coinedNum := int(gotData.(map[string]interface{})["multiply"].(float64)); coinedNum < 2 {
				params := make(url.Values)
				params.Add("bvid", videos[idx].Bvid)
				coinNum := min(2-coinedNum, reqCoinNum, user.Money)
				params.Add("multiply", strconv.Itoa(coinNum))
				params.Add("csrf", client.GetBiliJct())
				if t.config.(coinCfg).Like {
					params.Add("select_like", "1")
				}
				postCoinUrl := "http://api.bilibili.com/x/web-interface/coin/add"
				if err := client.CheckCode(client.PostForm(postCoinUrl, params)); err != nil {
					return err
				}
				reqCoinNum, user.Money = reqCoinNum-coinNum, user.Money-coinNum
			}
		}
		if reqCoinNum > 0 {
			if user.Money == 0 {
				return fmt.Errorf("用户硬币数量不足以完成投币任务")
			} else if len(idxs) == 0 {
				return fmt.Errorf("视频资源不足以完成投币任务")
			}
		}
		t.result = "投币成功"
		return nil
	},
}

type reward struct {
	IsWatch bool `mapstructure:"watch"`
	IsShare bool `mapstructure:"share"`
	Coin    int  `mapstructure:"coins"`
}

var rd *reward

func getReward() (*reward, error) {
	if rd != nil {
		return rd, nil
	}
	url := "http://api.bilibili.com/x/member/web/exp/reward"
	data, err := client.RecData(client.Get(url))
	if err != nil {
		return nil, err
	}
	rd = &reward{}
	err = mapstructure.Decode(data, rd)
	return rd, err
}

type video struct {
	Aid      int    `json:"aid"`
	Bvid     string `json:"bvid"`
	Cid      int    `json:"cid"`
	Duration int    `json:"duration"`
}

var videos []video

func getVideos(ps int) ([]video, error) {
	if videos != nil && len(videos) == ps {
		return videos, nil
	}
	url := "https://api.bilibili.com/x/web-interface/dynamic/region"
	params := [][]string{{"ps", strconv.Itoa(ps)}, {"rid", "1"}}
	data, err := client.RecData(client.GetWithParams(url, params))
	if err != nil {
		return nil, err
	}
	vMaps, tf := data.(map[string]interface{})["archives"].([]interface{})
	if !tf {
		return nil, fmt.Errorf("视频数据转换失败")
	}
	videos = make([]video, ps)
	for i, vMap := range vMaps {
		if err := mapstructure.Decode(vMap, &videos[i]); err != nil {
			return nil, err
		}
	}
	return videos, nil
}

func min(arr ...int) int {
	ret := math.MaxInt32
	for _, v := range arr {
		if ret > v {
			ret = v
		}
	}
	return ret
}
