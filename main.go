package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
	fc.StartHttp(HandleHttpRequest)
}

//func HandleHttpRequestDefault(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
//	w.WriteHeader(http.StatusOK)
//	w.Header().Add("Content-Type", "text/plain")
//	w.Write([]byte("hello, world!\n"))
//	return nil
//}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	time.Sleep(time.Millisecond * 50)

	bodyData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	bodyStr := string(bodyData)

	headerStr := req.Header.Get("my1")

	fmt.Printf("body content: %v, header content: %v", bodyStr, headerStr)

	w.Header().Add("Content-Type", "text/plain")
	if bodyStr != headerStr {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(bodyData)
	return nil
}
