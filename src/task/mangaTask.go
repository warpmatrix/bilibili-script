package task

import (
	"fmt"
	"main/src/client"
	"main/src/domain"
	log "main/src/logger"
	"net/http"
	"net/url"
)

var mangaCfg interface{}
var platform string

func init() {
	mangaCfg = domain.LoadCfg("manga")
	if mangaCfg == nil {
		log.Info("用户未设置漫画相关的配置")
		return
	}
	taskList = append(taskList, mangaClockIn)
}

func mangaDefaultInit(msg string) error {
	platform = "android"
	return fmt.Errorf(msg)
}

var mangaClockIn = &task{
	name: "漫画签到",
	init: func() error {
		var isString bool
		platform, isString = mangaCfg.(string)
		if !isString {
			return mangaDefaultInit("manga 配置的字段应为字符串")
		}
		if !(platform == "android" || platform == "ios") {
			return mangaDefaultInit("manga 配置的字段应为 android 或 ios")
		}
		return nil
	},
	impl: func(t *task) error {
		log.Info(fmt.Sprintf("【漫画签到平台】：%s", platform))
		params := make(url.Values)
		params.Add("platform", platform)
		url := "https://manga.bilibili.com/twirp/activity.v1.Activity/ClockIn"
		resp, err := client.PostForm(url, params)
		if err != nil {
			return err
		}
		switch resp.StatusCode {
		case http.StatusOK:
			t.result = "漫画签到成功"
		case http.StatusBadRequest:
			t.result = "请求错误，请检查今天是否已提前完成签到"
		default:
			return fmt.Errorf(resp.Status, "签到失败")
		}
		return nil
	},
}
