package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex

func test() {
	a := make(map[int]int)

	a[1] = 1
	a[2] = 3
	a[3] = 2
	a[4] = 5
	for i := 0; i < 2; i++ {

		go func(b map[int]int) {
			//fmt.Println("执行了")
			lock.Lock()
			fmt.Println("执行了")
			a[1] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	time.Sleep(time.Second)
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

}

func main() {
	test()
}
