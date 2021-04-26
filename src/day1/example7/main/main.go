package main

import (
	"day1/example7/add"
	"fmt"
)

func main() {
	var pip chan int
	pip = make(chan int, 1)
	go add.Add(100, 300, pip)
	sum := <-pip
	fmt.Println(sum)

}
