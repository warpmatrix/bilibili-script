package main

import (
	"log"
	"main/src/domain"
	"main/src/client"
)

func main() {
	blob, err := client.GetMsgBlob()
	if err != nil {
		log.Fatalln(err)
	}
	user, err := domain.ParseUser(blob)
	if err != nil {
		log.Fatalln(err)
	}
	user.PrintInfo()
}
