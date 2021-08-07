package task

import (
	"fmt"
	"main/src/client"
	"main/src/domain"
	log "main/src/logger"
	"net/http"
	"net/url"
)

// first level configuration name: manga
func init() {
	config := domain.LoadCfg("manga")
	if config == nil {
		log.Info("用户未设置漫画相关的配置")
		return
	}
	mangaClockIn.config = config
	taskList = append(taskList, mangaClockIn)
}

var mangaClockIn = &task{
	name: "漫画签到",
	initFunc: func(t *task) error {
		platform, isString := t.config.(string)
		if !isString {
			return fmt.Errorf("manga 配置的字段应为字符串")
		}
		if !(platform == "android" || platform == "ios") {
			return fmt.Errorf("manga 配置的字段应为 android 或 ios")
		}
		return nil
	},
	defaultCfg: "android",
	impl: func(t *task) error {
		platform := t.config.(string)
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
