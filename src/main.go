package main

import (
	"main/src/domain"
	"main/src/client"
	log "main/src/logger"
)

func main() {
	blob, err := client.GetMsgBlob()
	if err != nil {
		log.Fatal(err)
	}
	user, err := domain.ParseUser(blob)
	if err != nil {
		log.Fatal(err)
	}
	user.PrintInfo()
}
