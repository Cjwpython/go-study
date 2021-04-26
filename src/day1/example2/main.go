package main

// 函数引用
import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	a := 100
	b := 100
	sum := add(a, b)
	fmt.Println("sum:", sum)

	othersum := otheradd(a, b)
	fmt.Println("othersum:", othersum)
}
