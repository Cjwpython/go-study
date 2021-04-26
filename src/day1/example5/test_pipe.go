package main

import "fmt"

func test_pipe() {
	pipe := make(chan int, 3) // 创建一个管道，数据类型为int 容量为3
	pipe <- 1
	pipe <- 2
	pipe <- 3
	sum := <-pipe
	fmt.Println(sum)
	fmt.Println(len(pipe), cap(pipe))

}
