package task

import (
	"fmt"
	log "main/src/logger"
)

type task struct {
	name   string
	run    func() error
	result string
}

var taskList []*task

func RunTasks() {
	for _, task := range taskList {
		err := task.run()
		if err != nil {
			log.Error(err)
		} else {
			log.Info(fmt.Sprintf("【%s】:%s", task.name, task.result))
		}
	}
}
