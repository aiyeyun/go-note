package main

import (
	"net/http"
)

func main()  {

	// curl https://localhost:10086

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})

	http.ListenAndServeTLS(":10086", "/home/yeyun/WWW/golang/note/https/cert.pem", "/home/yeyun/WWW/golang/note/https/key.pem", nil)
}