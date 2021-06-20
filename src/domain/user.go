package domain

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	Level Level  `json:"level_info" mapstructure:"level_info"`
	Uname string `json:"uname" mapstructure:"uname"`
	// the number of coins
	Money int `json:"money" mapstructure:"money"`
}

type Level struct {
	CurExp   int `json:"current_exp" mapstructure:"current_exp"`
	CurLevel int `json:"current_level" mapstructure:"current_level"`
	// if the level is six, then nextExp is a string ("--")
	// else nextExp is an int
	NextExp interface{} `json:"next_exp" mapstructure:"next_exp"`
}

func ParseUser(blob []byte) (User, error) {
	user, msg := User{}, Message{}
	err := json.Unmarshal(blob, &msg)
	if err != nil {
		return user, err
	}
	switch msg.Code {
	case SUCCESS:
		err = mapstructure.Decode(msg.Data, &user)
	case NOT_LOGIN:
		err = fmt.Errorf("用户信息已过期，请重新绑定你的 cookie 信息")
	default:
		err = fmt.Errorf("转换用户信息失败")
	}
	return user, err
}

func (user User) PrintInfo() {
	log.Println("【用户名】:", user.Uname)
	log.Println("【硬币数量】:", user.Money)
	log.Println("【当前等级】:", user.Level.CurLevel)
	log.Println("【当前经验】:", user.Level.CurExp)
	switch t := user.Level.NextExp.(type) {
	case float64:
		log.Println("【距离下一级的经验】：", t)
	case string:
		log.Println("【距离一下级的经验】：当前已经是最高级")
	default:
		log.Println("【距离一下级的经验】：类型转换错误")
	}
}
