package main

import "fmt"

func main()  {
	Go()
}

func Go()  {

	// pannic 后面的代码就不会执行了
	// pannic 后的错误信息 可以在 defer 里用 recover 来接收 ,再处理 pannic
	// 多个defer 都会执行

	defer func() {
		fmt.Println("end")
		if err := recover(); err != nil {
			fmt.Println("error info:", err)
		}
	}()

	defer fmt.Println(" i love defer ")

	fmt.Println("start ...")
	panic(" error info : xxx ")

	fmt.Println("center")
}