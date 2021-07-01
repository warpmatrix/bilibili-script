package task

import (
	"encoding/json"
	"fmt"
	"main/src/client"
	"main/src/domain"
	log "main/src/logger"
	"net/url"
)

type MangaCfg struct {
	platform string
}

var mangaCfg MangaCfg

func init() {
	m := domain.LoadCfg("manga")
	if m == nil {
		log.Info("用户未设置漫画相关的配置")
		return
	}
	if task := initMangaTask(m); task != nil {
		taskList = append(taskList, task)
	}
}

func initMangaTask(m interface{}) *task {
	var isString bool
	mangaCfg.platform, isString = m.(string)
	if !isString {
		log.Warning("manga 配置的字段应为字符串")
		return nil
	}
	var mangaTask *task
	switch mangaCfg.platform {
	case "android":
		fallthrough
	case "ios":
		mangaTask = &task{
			name: "漫画签到",
		}
		mangaTask.run = mangaTask.mangaSignIn
	default:
		log.Warning("manga 配置的字段应为 android 或 ios")
	}
	return mangaTask
}

func (t *task) mangaSignIn() error {
	log.Info(fmt.Sprintf("【漫画签到平台】：%s", mangaCfg.platform))
	param := make(url.Values)
	param.Add("platform", mangaCfg.platform)
	url := "https://manga.bilibili.com/twirp/activity.v1.Activity/ClockIn"
	blob, err := client.PostForm(url, param)
	if err != nil {
		return err
	}
	var msg domain.Message
	err = json.Unmarshal(blob, &msg)
	if err != nil {
		return err
	}
	switch msg.Code {
	case 0:
		t.result = "漫画签到成功"
	default:
		t.result = "今天已经完成签到了"
	}
	return nil
}
