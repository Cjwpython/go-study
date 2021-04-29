package main

import "fmt"

func test(a int, b ...int) {
	fmt.Println(a)
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(len(b))
}

func concat(a string, arg ...string) (result string) {
	result = a

	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}
	return
}
func main() {
	var a int = 1
	var b int = 2
	var c int = 3
	test(a, b, c)

	str := concat("hello", " ", "world")
	fmt.Println(str)

}
