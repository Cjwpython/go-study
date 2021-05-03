package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), "-")
	fmt.Print(f(20), "-")
	fmt.Print(f(300))

}

func Adder() func(int) int {
	var x int
	defer fmt.Println("func end")
	fmt.Println("func start ")
	return func(delta int) int {
		x += delta
		return x
	}
}
