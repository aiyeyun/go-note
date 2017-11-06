package main

import (
	"runtime"
	"fmt"
	"time"
)

var (
	a_chan chan string = make(chan string, 1)
	b_chan chan string = make(chan string, 1)
	c_chan chan string = make(chan string, 1)
	d_chan chan string = make(chan string, 1)
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go a()
	go b()
	go c()
	go d()

	fmt.Println(<- a_chan)
	fmt.Println(<- b_chan)
	fmt.Println(<- c_chan)
	fmt.Println(<- d_chan)
}

func a()  {
	fmt.Println("A 正在执行...", time.Now())
	time.Sleep(time.Second * 5)
	a_chan <- "A 执行完毕 " + time.Now().String()
}

func b()  {
	fmt.Println("B 正在执行...", time.Now())
	time.Sleep(time.Second * 5)
	b_chan <- "B 执行完毕 " + time.Now().String()
}

func c()  {
	fmt.Println("C 正在执行...", time.Now())
	time.Sleep(time.Second * 5)
	c_chan <- "C 执行完毕 " + time.Now().String()
}

func d()  {
	fmt.Println("D 正在执行...", time.Now())
	time.Sleep(time.Second * 5)
	d_chan <- "D 执行完毕" + time.Now().String()
}