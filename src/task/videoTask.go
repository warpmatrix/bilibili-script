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

var videoCfg map[string]interface{}

// first level configuration name: video
func init() {
	videoCfg = domain.LoadCfg("video").(map[string]interface{})
	if videoCfg == nil {
		log.Info("用户未配置视频相关的配置")
		return
	}
	// second level configuration name: daily
	for _, daily := range videoCfg["daily"].([]interface{}) {
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
	if videoCfg["coin"] != nil {
		taskList = append(taskList, coinVideo)
	}
}

var watchVideo = &task{
	name: "模拟观看视频",
	init: defaultInit,
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

var shareVideo = &task{
	name: "分享视频",
	init: defaultInit,
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
		params.Add("csrf", domain.GetBiliJct())
		url := "https://api.bilibili.com/x/web-interface/share/add"
		blob, err := client.ParseResp(client.PostForm(url, params))
		if err != nil {
			return err
		}
		if _, err := domain.CheckMsgBlob(blob); err != nil {
			return err
		}
		t.result = "成功模拟分享视频"
		return nil
	},
}

var coinVideo = &task{
	name: "视频投币",
	init: defaultInit,
	impl: func(t *task) error {
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
	videos = make([]video, ps)
	for i, vMap := range vMaps {
		if err := mapstructure.Decode(vMap, &videos[i]); err != nil {
			return nil, err
		}
	}
	return videos, nil
}
