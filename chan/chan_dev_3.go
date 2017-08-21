package main

import "fmt"

//import "fmt"

func main() {
	ch := make(chan map[string]interface{}, 3)
	go Goa(ch)
	go Gob(ch)
	go Goc(ch)

	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}

func Goa(ch chan map[string]interface{})  {
	obj := make(map[string]interface{})
	obj["a"] = "你管我是啥类型"
	fmt.Println("a over")
	ch <- obj
}

func Gob(ch chan map[string]interface{})  {
	obj := make(map[string]interface{})
	subObj := make(map[string]string)
	subObj["b_b"] = "b_b"
	subObj["bb_bb"] = "bb_bb"
	obj["b"] = subObj
	fmt.Println("b over")
	ch <- obj
}

func Goc(ch chan map[string]interface{})  {
	obj := make(map[string]interface{})
	obj["c"] = "你管我是啥类型"
	fmt.Println("c over")
	ch <- obj
}