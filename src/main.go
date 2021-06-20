package main

import (
	"log"
	"main/src/domain"
)

func main() {
	blob, err := getMsgBlob()
	if err != nil {
		log.Fatalln(err)
	}
	user, err := domain.ParseUser(blob)
	if err != nil {
		log.Fatalln(err)
	}
	user.PrintInfo()
}
