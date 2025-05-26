package main

import (
	"fmt"
	"strconv"
)

//型変換

func main() {
	var s string = "100"
	fmt.Printf("s = %T\n", s)
	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

}
