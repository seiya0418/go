package main

import "fmt"

//関数

func Plus (a int, b int) int {
	return a + b
}

//複数の戻り値
func Div(a int, b int) (int, int) {
	return a / b, a % b

}

//返り値
func Double(a int) (int, int) {
	return a * 2, a * 3
}

//返り値のない関数
func NoReturn() {
	fmt.Println("No Return")
}

func main() {
	fmt.Println(Plus(1, 2))
	fmt.Println(Div(10, 3))
	fmt.Println(Double(2))
	NoReturn()
}

