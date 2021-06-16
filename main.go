package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	// fmt.Println(os.Getenv("BILI_JCT"))
	// fmt.Println(os.Getenv("DEDEUSERID"))
	// fmt.Println(os.Getenv("SESSDATA"))
	// url := "https://www.baidu.com/"
	url := "https://github.com/"
	tr := &http.Transport{
		TLSClientConfig:	&tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("error: status code ", res.StatusCode)
		return
	}
	fmt.Println(res.Status)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body))
}
