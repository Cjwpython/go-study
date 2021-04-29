package main

import "fmt"

func main() {
	var a = 123
	switch a {
	case 0:
		fallthrough
	case 123:
		fmt.Println(123123123)
	case 1231:
		fmt.Println(1231231)
	default:
		fmt.Println("asdasdas")
	}
}
