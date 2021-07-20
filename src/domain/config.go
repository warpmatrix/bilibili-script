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
		if !filepath.IsAbs(filePath) {
			workPath, _ := os.Getwd()
			filePath = filepath.Join(workPath, filePath)
		}
		log.Info("用户未在定义路径", filePath, "下设置自定义配置，使用默认配置")
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
