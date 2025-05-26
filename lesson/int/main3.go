package main

import "fmt"

func main() {
	var i int = 100
	fmt.Println(i + 50)

	var i2 int64 = 200
	fmt.Println("%T\n", i2)
	fmt.Println("%T\n", int32(i2))
}