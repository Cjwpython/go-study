package main

import (
	"day1/example6/calc"
	"fmt"
)

func main() {

	sum := calc.Add(1000, 5000)
	fmt.Println(sum)
	sub := calc.Sub(1000, 5000)
	fmt.Println(sub)
}
