package task

import (
	"fmt"
	log "main/src/logger"
)

type task struct {
	name   string
	result string
	impl   func(*task) error

	config     interface{}
	initFunc   func(*task) error
	checkFuncs map[string]func(interface{}) error
	defaultCfg interface{}
	cfgMap     map[string]interface{}
}

var taskList []*task

func RunTasks() {
	for _, task := range taskList {
		if err := task.init(); err != nil {
			log.Warning(err)
			task.config = task.defaultCfg
		}
		if err := task.run(); err != nil {
			log.Error(fmt.Sprintf("【%s】：%s", task.name, err))
		} else {
			log.Info(fmt.Sprintf("【%s】：%s", task.name, task.result))
		}
	}
}

func (t *task) run() error {
	return t.impl(t)
}

func (t *task) init() error {
	if t.initFunc != nil {
		return t.initFunc(t)
	}
	return nil
}

func (t *task) setCfg(key string, val interface{}) error {
	if f := t.checkFuncs[key]; f != nil {
		if err := f(val); err != nil {
			return fmt.Errorf("%s，使用默认配置字段", err)
		} else {
			t.cfgMap[key] = val
		}
	} else {
		return fmt.Errorf("未能识别字段：%s", key)
	}
	return nil
}
