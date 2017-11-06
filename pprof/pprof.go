package main

import (
	"time"
	"net/http"
	"runtime/pprof"
)

var quit chan struct{} = make(chan struct{})

func f() {
	<-quit
}

func main() {
	for i := 0; i < 10000; i++ {
		go f()
	}

	go goppf() //启用跟踪查看
	for {
		time.Sleep(1 * time.Second)
	}
}

func goppf() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":10086", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}