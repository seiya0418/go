package main

import "fmt"

//型
//文字列型

func main() {
	var s string = "Hello"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Println(si)

	fmt.Println(`test
	test
	    	test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(s[0])
	
}