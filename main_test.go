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
	resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		t.Error("curl failed")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("read body failed")
	}

	// TODO 这里不优雅
	fmt.Println(string(body) == `
  __  _  _ ____ ____  __  _  _ ____    __ _  __ ____ ____ ____     __  ____ ____ 
 / _\/ )( (  __) ___)/  \( \/ |  __)  (  ( \/  (_  _|  __) ___)   / _\(  _ (  _ \
/    \ /\ /) _)\___ (  O ) \/ \) _)   /    (  O ))(  ) _)\___ \  /    \) __/) __/
\_/\_(_/\_|____|____/\__/\_)(_(____)  \_)__)\__/(__)(____|____/  \_/\_(__) (__)

`)
}
