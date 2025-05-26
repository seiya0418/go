package main

import "fmt"

//switch
//型スイッチ

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("Hello")
	anything(1)

	var x interface{} = 3
	i, isInt := x.(int)
	fmt.Println(i, isInt)

	f, isFloat := x.(float64)
	fmt.Println(f, isFloat)

	s, isString := x.(string)
	fmt.Println(s, isString)
}