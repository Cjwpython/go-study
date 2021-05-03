package main

import "fmt"

func main() {
	var i int
	fmt.Println(i)
	j := new(int)
	*j = 123
	fmt.Println(*j)
}
