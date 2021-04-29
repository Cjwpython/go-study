package main

import "fmt"

type add_func func(int, int) int

func add(a int, b int) int {
	return a + b
}
func add_num(op add_func, a int, b int) int {
	return op(a, b)
}
func main() {
	c := add
	fmt.Println(add)
	sum := add_num(c, 10, 20)
	fmt.Println(sum)

}
