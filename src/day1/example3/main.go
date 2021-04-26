package main

import (
	"time"
)

// go route 演示

func main() {
	a := 100
	b := 100
	go testsum(a, b)
	time.Sleep(time.Second)

}
