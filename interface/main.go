package main

import "fmt"

func main()  {
	var inter interface{}
	var val interface{} = (*interface{})(nil)
	val = inter
	fmt.Println(val)
}
