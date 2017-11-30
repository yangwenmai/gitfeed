package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetJSON get url json数据
func GetJSON(url string) (resp *http.Response, data string, err error) {
	resp, err = http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bs, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("http get error, status = %d, resp = %v", resp.StatusCode, string(bs))
		data = string(bs)
		return
	}
	bs, err := ioutil.ReadAll(resp.Body)
	data = string(bs)
	if err != nil {
		return
	}

	return
}
