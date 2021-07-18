package main

import (
	"main/src/client"
	"main/src/domain"
	log "main/src/logger"
	"main/src/task"
)

func main() {
	user, err := getUserInfo()
	if err != nil {
		log.Fatal(err)
	}
	user.PrintInfo()
	task.RunTasks()
}

func getUserInfo() (domain.User, error) {
	url := "https://api.bilibili.com/x/web-interface/nav"
	blob, err := client.ParseResp(client.Get(url))
	if err != nil {
		return domain.User{}, err
	}
	return domain.ParseUser(blob)
}
