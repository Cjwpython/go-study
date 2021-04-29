package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum string

	fmt.Scanf("%s", &sum)
	a, err := strconv.Atoi(sum)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(a)
	}

}
