package main

import (
	"fmt"
)

func main()  {
	ch := make([]chan bool, 10)

	for i:=0; i<10; i++ {
		ch[i] = make(chan bool)
		go Go(i, ch[i])
	}

	for _, v := range ch {
		<- v
		//fmt.Println(<- v)
	}

}

func Go(index int, ch chan bool)  {
	if index == 9 {
	}
	fmt.Println("index:", index)
	ch <- true
}