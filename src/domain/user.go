package domain

import (
	"github.com/mitchellh/mapstructure"
	log "main/src/logger"
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

func GetUserInfo(blob []byte) (*User, error) {
	data, err := ParseBlob(blob)
	if err != nil {
		return nil, err
	}
	var user User
	err = mapstructure.Decode(data, &user)
	return &user, err
}

func (user *User) PrintInfo() {
	log.Info("【用户名】:", user.Uname)
	log.Info("【硬币数量】:", user.Money)
	log.Info("【当前等级】:", user.Level.CurLevel)
	log.Info("【当前经验】:", user.Level.CurExp)
	switch t := user.Level.NextExp.(type) {
	case float64:
		log.Info("【距离下一级的经验】：", t)
	case string:
		log.Info("【距离一下级的经验】：当前已经是最高级")
	default:
		log.Info("【距离一下级的经验】：类型转换错误")
	}
}
