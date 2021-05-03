package main

import "fmt"

func deractor(a func()) {
	fmt.Println("func: start")
	a()
	fmt.Println("func: start")
}

func test() {
	fmt.Println("hello world")
}
func main() {
	var a = test
	deractor(a)
}
