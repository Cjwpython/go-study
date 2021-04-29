package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i < 6; i++ {
		str := strings.Repeat("A", i)
		fmt.Println(str)
	}
}
