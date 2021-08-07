package main

import (
	"main/src/domain"
	log "main/src/logger"
	"main/src/task"
)

func main() {
	user, err := domain.GetUserInfo()
	if err != nil {
		log.Fatal(err)
	}
	user.PrintInfo()
	task.RunTasks()
}
