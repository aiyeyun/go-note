package main

import "fmt"

func main()  {

	// 创建有缓存的 channel , chan 未被填满前 是不会阻塞的,
	// 创建无缓存的 channel , chan 读 or 写 都会被阻塞的 ,

	// 这是一个 有缓存的 channel

	ch := make(chan bool, 2)
	go GoA(ch)
	go GoB(ch)

	for i:=0; i<2; i++ {
		<- ch
	}
}

func GoA(ch chan bool)  {
	fmt.Println("A执行了")
	ch <- true
}

func GoB(ch chan bool)  {
	fmt.Println("B执行了")
	ch <- true
}