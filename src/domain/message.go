package domain

import (
	"fmt"
	"os"
)

type Message struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS   = 0
	NOT_LOGIN = -101
)

func GetCookie() string {
	return fmt.Sprintf("bili_jct=%s;SESSDATA=%s;DedeUserID=%s;", os.Getenv("BILI_JCT"), os.Getenv("SESSDATA"), os.Getenv("DEDEUSERID"))
}
