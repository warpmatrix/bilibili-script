package task

import (
	"fmt"
	log "main/src/logger"
)

type task struct {
	name        string
	result      string
	init        func() error
	defaultInit func()
	impl        func(*task) error
}

var taskList []*task

func RunTasks() {
	for _, task := range taskList {
		if err := task.init(); err != nil {
			log.Warning(err)
			task.defaultInit()
		}
		if err := task.run(); err != nil {
			log.Error(err)
		} else {
			log.Info(fmt.Sprintf("【%s】：%s", task.name, task.result))
		}
	}
}

func (t *task) run() error {
	return t.impl(t)
}

var defaultInit = func() error { return nil }
