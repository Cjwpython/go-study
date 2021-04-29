package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var input string
	rand.Seed(time.Now().Unix())
	var n int
	n = rand.Intn(100)

	fmt.Scanf("%s\n", &input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("error")

	}
	fmt.Println(n)
	switch num {
	case n:
		fmt.Println("sucess")
	default:
		fmt.Println("faild")
	}
}
