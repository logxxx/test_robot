package main

import (
	"fmt"
	"net/http"
	"test_robot/utils/log"
)

func getReply(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nihao")
}

func main() {
	http.HandleFunc("/robot/get_reply", getReply)

	addr := "127.0.0.1:8000"
	log.Infof("监听地址:%v", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Errorf("ListenAndServe err:%v", err)
	}
}