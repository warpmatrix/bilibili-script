package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// url := "https://t.bilibili.com/"
	// url := "https://account.bilibili.com/account/home"
	url := "https://api.bilibili.com/x/web-interface/nav"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := CreateGetRequest(url)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalln(res.Status)
		return
	}
	fmt.Println(url, res.Status)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body))
}
