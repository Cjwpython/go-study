package main

import (
	"time"
)

// go route 演示

func main() {
	for i := 0; i < 100; i++ {
		go print_int(i)
	}
	time.Sleep(time.Second * 4)

}
