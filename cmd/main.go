package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Println("读取告警通知信息失败, err=", err.Error())
		}
		fmt.Println("读取告警通知信息：", string(body))
		fmt.Fprintf(writer, "ok")
	})
	fmt.Println("webhook starting ------------------------------------------------")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		panic(err)
	}
}
