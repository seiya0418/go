package main

import "fmt"

//関数
//クロージャー

func Later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		s = s + next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("World"))
	fmt.Println(f("!"))
}