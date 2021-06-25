package domain

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var cfg map[string]interface{}

const (
	filename  = "config.yaml"
	directory = "./"
)

func init() {
	filePath := filepath.Join(directory, filename)
	blob, err := os.ReadFile(filePath)
	if err != nil {
		goto handleErr
	}
	err = yaml.Unmarshal(blob, &cfg)
	if err != nil {
		goto handleErr
	}
	return
handleErr:
	log.Println(err)
}

func LoadCfg(name string) interface{} {
	return cfg[name]
}
