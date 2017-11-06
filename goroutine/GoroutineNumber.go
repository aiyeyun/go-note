package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		fmt.Println(111)
	}()
	go func() {
		fmt.Println(111)
	}()

	fmt.Println(runtime.NumGoroutine())
}
