package domain

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
	SUCCESS   = 0
	NOT_LOGIN = -101
	CSRF_FAIL = -111
	REQ_FAIL = -400
	NO_VIDEO = 10003
)

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
			switch kv[0] {
			case "bili_jct":
				bili_jct = kv[1]
			case "SESSDATA":
				sessdata = kv[1]
			case "DedeUserID":
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

func ParseBlob(blob []byte) (interface{}, error) {
	msg, err := CheckMsgBlob(blob)
	if err != nil {
		return nil, err
	}
	return msg.Data, nil
}

func CheckMsgBlob(blob []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(blob, &msg)
	if err != nil {
		return nil, err
	}
	switch msg.Code {
	case SUCCESS:
	case NOT_LOGIN:
		return nil, fmt.Errorf("用户信息已过期，请重新绑定你的 cookie 信息：%s", string(blob))
	case CSRF_FAIL:
		return nil, fmt.Errorf("用户 bili_jct 信息错误：%s", string(blob))
	case REQ_FAIL:
		return nil, fmt.Errorf("请求错误：%s", string(blob))
	case NO_VIDEO:
		return nil, fmt.Errorf("不存在该稿件：%s", string(blob))
	default:
		return nil, fmt.Errorf(string(blob))
	}
	return &msg, nil
}
