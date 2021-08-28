package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RecData(resp *http.Response, err error) (interface{}, error) {
	blob, err := ParseResp(resp, err)
	if err != nil {
		return nil, err
	}
	return ParseBlob(blob)
}

func CheckCode(resp *http.Response, err error) error {
	blob, err := ParseResp(resp, err)
	if err != nil {
		return err
	}
	_, err = ParseBlob(blob)
	return err
}

func ParseBlob(blob []byte) (interface{}, error) {
	msg, err := checkMsgBlob(blob)
	if err != nil {
		return nil, err
	}
	return msg.Data, nil
}

func ParseResp(resp *http.Response, err error) ([]byte, error) {
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
