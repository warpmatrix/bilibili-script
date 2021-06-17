package main

import (
	"fmt"
	"os"
	"net/http"
)

func CreateGetRequest(url string) (*http.Request, error)  {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36")
	req.Header.Set("Cookie", getCookie())
	return req, err
}

func getCookie() string {
	return fmt.Sprintf("bili_jct=%s;SESSDATA=%s;DedeUserID=%s;", os.Getenv("BILI_JCT"), os.Getenv("SESSDATA"), os.Getenv("DEDEUSERID"))
}