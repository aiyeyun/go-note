package main

import "fmt"

func main()  {

	// 创建有缓存的 channel , chan 未被填满前 是不会阻塞的,
	// 创建无缓存的 channel , chan 读 or 写 都会被阻塞的 ,

	//这是一个 无缓存的 channel

	ch := make(chan int)
	go goA(ch)
	go goB(ch)

	<- ch
	<- ch

	fmt.Println("我执行完了")
}

func goA(ch chan int)  {
	fmt.Println("A")
	ch <- 1
}

func goB(ch chan int)  {
	fmt.Println("B")
	ch <- 2
}
