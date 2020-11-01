package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_hello(t *testing.T) {
	go main()
	time.Sleep(time.Second) // 给main函数续一秒，确保main在http.Get之前执行
	resp, err := http.Get("http://localhost:8010/")
	if err != nil {
		t.Error("curl failed")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("read body failed")
	}

	// TODO 这里不优雅
	fmt.Println(string(body) == "This is a awesome note app!")
}
