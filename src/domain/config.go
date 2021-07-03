package domain

import (
	log "main/src/logger"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var cfg map[string]interface{}

const (
	filename  = "config.yaml"
	directory = "config/"
)

func init() {
	filePath := filepath.Join(directory, filename)
	blob, err := os.ReadFile(filePath)
	if err != nil {
		log.Info("用户未进行自定义设置，使用默认配置")
		return
	}
	err = yaml.Unmarshal(blob, &cfg)
	if err != nil {
		log.Warning("读取用户配置失败：", err, "（使用默认配置）")
	}
}

func LoadCfg(name string) interface{} {
	return cfg[name]
}
