package main

import "fmt"

//for


func main() {
	/*
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}
	*/

	/*
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}*/

	/*
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}
	*/

	/*
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	*/

	/*
	sl := []string{"A", "B", "C"}
	for _, v := range sl {
		fmt.Println(v)
	}
	*/

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
