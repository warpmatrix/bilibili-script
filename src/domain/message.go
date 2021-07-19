package domain

import (
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
)

var cookie_kv string

func getCookie() string {
	bili_jct := os.Getenv("BILI_JCT")
	sessdata := os.Getenv("SESSDATA")
	dedeuserid := os.Getenv("DEDEUSERID")
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
	return fmt.Sprintf("bili_jct=%s;SESSDATA=%s;DedeUserID=%s;", bili_jct, sessdata, dedeuserid)
}

func GetCookie() string {
	if len(cookie_kv) == 0 {
		cookie_kv = getCookie()
	}
	return cookie_kv
}
