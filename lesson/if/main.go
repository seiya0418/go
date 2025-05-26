package main

import "fmt"

//if

func main() {
	var x int = 10
	if x > 0 {
		fmt.Println("x is positive")
	}else if x < 0 {
		fmt.Println("x is negative")
	}else {
		fmt.Println("x is zero")
	}
}