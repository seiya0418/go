package main

import (
	"fmt"
	"strconv"
)

//エラーハンドリング

func main() {
	var s string = "100"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)
}