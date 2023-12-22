package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("正在启动Web服务器...")
	const WebBindAddress string = "0.0.0.0:8080"
	http.HandleFunc("/", index)
	err := http.ListenAndServe(WebBindAddress, nil)
	if err != nil {
		log.Print(err)
		return
	}
}

func index(res http.ResponseWriter, req *http.Request) {

	log.Print("请求了 API on [index]")
}
