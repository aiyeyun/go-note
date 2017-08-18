package main

import (
	"net/http"
	"time"
	_ "net/http/pprof"
)

func main()  {
	go func() {
		http.ListenAndServe("0.0.0.0:10086", nil)
	}()

	time.Sleep(time.Second * 10000)
}
