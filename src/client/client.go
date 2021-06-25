package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"main/src/domain"
	"net/http"
)

func CreateGetRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36")
	req.Header.Set("Cookie", domain.GetCookie())
	return req, err
}

func GetMsgBlob() ([]byte, error) {
	// url := "https://t.bilibili.com/"
	// url := "https://account.bilibili.com/account/home"
	url := "https://api.bilibili.com/x/web-interface/nav"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := CreateGetRequest(url)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(res.Status)
	}
	ret, err := ioutil.ReadAll(res.Body)
	return ret, err
}
