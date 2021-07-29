package task

import (
	"fmt"
	"main/src/client"
	"main/src/domain"
	log "main/src/logger"
	"math/rand"
	"net/url"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

type VideoCfg struct {
	Dailies []string `mapstructure:"daily"`
	Coin    int      `mapstructure:"coin"`
}

var videoCfg VideoCfg

func init() {
	m := domain.LoadCfg("video")
	if m == nil {
		log.Info("用户未配置视频相关的配置")
		return
	}
	err := mapstructure.Decode(m, &videoCfg)
	if err != nil {
		log.Warning("视频相关配置读取失败", err)
	}
	for _, daily := range videoCfg.Dailies {
		switch daily {
		case "watch":
			taskList = append(taskList, watchVideo)
		// case "share":
		// 	taskList = append(taskList, shareVideo)
		default:
		}
	}
	// if videoCfg.Coin > 0 {
	// 	taskList = append(taskList, coinVideo)
	// }
}

type reward struct {
	IsWatch bool `mapstructure:"watch"`
	IsShare bool `mapstructure:"share"`
	Coin    int  `mapstructure:"coins"`
}

var rd *reward

type video struct {
	Aid      int `json:"aid"`
	Cid      int `json:"cid"`
	Duration int `json:"duration"`
}

var watchVideo = &task{
	name: "模拟观看视频",
	init: defaultInit,
	impl: func(t *task) error {
		rd, err := getReward()
		if err != nil {
			return fmt.Errorf("获取用户信息失败：%s", err)
		}
		if rd.IsWatch {
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
		params.Add("csrf", domain.GetBiliJct())
		url := "https://api.bilibili.com/x/v2/history/report"
		blob, err := client.ParseResp(client.PostForm(url, params))
		if err != nil {
			return err
		}
		if _, err := domain.CheckMsgBlob(blob); err != nil {
			return err
		}
		t.result = "成功模拟观看视频"
		return nil
	},
}

func getReward() (*reward, error) {
	if rd != nil {
		return rd, nil
	}
	url := "http://api.bilibili.com/x/member/web/exp/reward"
	blob, err := client.ParseResp(client.Get(url))
	if err != nil {
		return nil, err
	}
	data, err := domain.ParseBlob(blob)
	if err != nil {
		return nil, err
	}
	rd = &reward{}
	err = mapstructure.Decode(data, rd)
	return rd, err
}

func getVideos(ps int) ([]video, error) {
	url := "https://api.bilibili.com/x/web-interface/dynamic/region"
	params := [][]string{{"ps", strconv.Itoa(ps)}, {"rid", "1"}}
	blob, err := client.ParseResp(client.GetWithParams(url, params))
	if err != nil {
		return nil, err
	}
	data, err := domain.ParseBlob(blob)
	if err != nil {
		return nil, err
	}
	vMaps, tf := data.(map[string]interface{})["archives"].([]interface{})
	if !tf {
		return nil, fmt.Errorf("视频数据转换失败")
	}
	ret := make([]video, ps)
	for i, vMap := range vMaps {
		if err := mapstructure.Decode(vMap, &ret[i]); err != nil {
			return nil, err
		}
	}
	return ret, nil
}
