package main

import (
	"time"
	"fmt"
)

func main() {

	var ch chan string = make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch <- "write"
	}()

	select {
	case val := <- ch :
		fmt.Println(val)
	case <- time.After(time.Second * 2):
		fmt.Println("超时了")
	}

}
