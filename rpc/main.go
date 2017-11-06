package main

import (
	"net/http"
	"net/rpc"
	"io"
	"fmt"
	"net"
)

type Watcher int

func (w *Watcher) GetInfo(arg int, result *int) error {
	*result = 10
	return nil
}

func main() {
	http.HandleFunc("/", demo)
	watcher := new(Watcher)
	rpc.Register(watcher)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("监听失败，端口可能已经被占用")
	}
	fmt.Println("正在监听10086端口")
	http.Serve(l, nil)
}

func demo(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "<html><body>ghj1976-123</body></html>")
}