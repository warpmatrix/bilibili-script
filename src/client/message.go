package client

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Message struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS      = 0
	NOT_LOGIN    = -101
	CSRF_FAIL    = -111
	REQ_FAIL     = -400
	NO_VIDEO     = 10003
	COIN_SELF    = 34002
	INVALID_COIN = 34003
	COIN_SHORT   = 34004
	OVER_COIN    = 34005
)

var codeText = map[int]string{
	NOT_LOGIN:    "用户信息已过期，请重新绑定你的 cookie 信息",
	CSRF_FAIL:    "用户 bili_jct 信息错误",
	REQ_FAIL:     "请求错误",
	NO_VIDEO:     "不存在该稿件",
	COIN_SELF:    "不能给自己的稿件投币",
	INVALID_COIN: "非法投币数目",
	COIN_SHORT:   "投币间隔过短",
	OVER_COIN:    "超过投币上限",
}

var cookie_kv, bili_jct, sessdata, dedeuserid string

func getCookie() string {
	bili_jct = os.Getenv("BILI_JCT")
	sessdata = os.Getenv("SESSDATA")
	dedeuserid = os.Getenv("DEDEUSERID")
	// use cookie first
	if cookies := os.Getenv("COOKIE"); len(cookies) > 0 {
		lines := strings.Split(cookies, "\n")
		for _, line := range lines {
			kv := strings.Split(line, "=")
			switch strings.ToLower(kv[0]) {
			case "bili_jct":
				bili_jct = kv[1]
			case "sessdata":
				sessdata = kv[1]
			case "dedeuserid":
				dedeuserid = kv[1]
			default:
			}
		}
	}
	cookie_kv = fmt.Sprintf("bili_jct=%s;SESSDATA=%s;DedeUserID=%s;", bili_jct, sessdata, dedeuserid)
	return cookie_kv
}

func GetCookie() string {
	if len(cookie_kv) == 0 {
		return getCookie()
	}
	return cookie_kv
}

func GetBiliJct() string {
	if len(bili_jct) == 0 {
		getCookie()
	}
	return bili_jct
}

func checkMsgBlob(blob []byte) (*Message, error) {
	var msg Message
	if err := json.Unmarshal(blob, &msg); err != nil {
		return nil, err
	}
	if msg.Code != SUCCESS {
		if len(codeText[msg.Code]) > 0 {
			return nil, fmt.Errorf("%s：%s", codeText[msg.Code], string(blob))
		} else {
			return nil, fmt.Errorf("%s", string(blob))
		}
	}
	return &msg, nil
}
