package main

import "fmt"

func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

func main() {
	n := 2
	for i := 0; i <= n; i++ {
		s := fab(i)
		fmt.Println(s)
	}
}
