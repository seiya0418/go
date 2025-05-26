package main

import "fmt"

//関数
//無名関数

func main() {
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	f := func(x, y int) {
		fmt.Println(x + y)
	}
	f(10, 20)
}