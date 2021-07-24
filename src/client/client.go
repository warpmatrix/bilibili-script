package client

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"main/src/domain"
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
	req.Header.Set("Cookie", domain.GetCookie())
	req.Header.Add("referer", "https://www.bilibili.com/")
	req.Header.Add("connection", "keep-alive")
	return req, err
}

func Get(url string) ([]byte, error) {
	req, err := newRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return do(req)
}

func Post(url string, contentType string, body io.Reader) ([]byte, error) {
	req, err := newRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("charset", "UTF-8")
	return do(req)
}

func PostJson(url string, jsonBlob []byte) ([]byte, error) {
	return Post(url, "application/json", strings.NewReader(string(jsonBlob)))
}

func PostForm(url string, data url.Values) ([]byte, error) {
	return Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func do(req *http.Request) ([]byte, error) {
	wait(1, 3)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	ret, err := ioutil.ReadAll(resp.Body)
	return ret, err
}

func wait(minSec, maxSec int) {
	waitTime := time.Duration(rand.Intn(maxSec-minSec) + minSec)
	<-time.After(waitTime * time.Second)
}
