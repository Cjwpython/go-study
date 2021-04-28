package main

import (
	"fmt"
)

func print_list(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(i, n-i)
	}
}

func main() {
	n := 4
	print_list(n)
}
