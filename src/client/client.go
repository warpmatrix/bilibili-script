package client

import (
	"crypto/tls"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client *http.Client

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
}

func newRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.101 Safari/537.36")
	req.Header.Set("Cookie", GetCookie())
	req.Header.Add("referer", "https://www.bilibili.com/")
	req.Header.Add("connection", "keep-alive")
	return req, err
}

func get(url string, params [][]string) (*http.Response, error) {
	req, err := newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	for _, param := range params {
		query.Add(param[0], param[1])
	}
	req.URL.RawQuery = query.Encode()
	return do(req)
}

func Get(url string) (*http.Response, error) {
	return get(url, nil)
}

func GetWithParams(url string, params[][]string) (*http.Response, error) {
	return get(url, params)
}

func post(url string, contentType string, body io.Reader) (*http.Response, error) {
	req, err := newRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("charset", "UTF-8")
	return do(req)
}

func PostJson(url string, jsonBlob []byte) (*http.Response, error) {
	return post(url, "application/json", strings.NewReader(string(jsonBlob)))
}

func PostForm(url string, params url.Values) (*http.Response, error) {
	return post(url, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
}

func do(req *http.Request) (*http.Response, error) {
	wait(1, 3)
	return client.Do(req)
}

func wait(minSec, maxSec int) {
	waitTime := time.Duration(rand.Intn(maxSec-minSec) + minSec)
	<-time.After(waitTime * time.Second)
}
