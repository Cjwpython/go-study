package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//var rwlock sync.Mutex
var rwlock sync.RWMutex

func testRWlock() {
	a := make(map[int]int)
	var count int32
	a[1] = 1
	a[2] = 3
	a[3] = 2
	a[4] = 5
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			b[1] = rand.Intn(100)
			time.Sleep(time.Millisecond * 10)
			rwlock.Unlock()

		}(a)
	}

	for i := 0; i < 100; i++ {

		go func(b map[int]int) {
			for {
				rwlock.Lock()
				time.Sleep(time.Millisecond)
				rwlock.Unlock()
				atomic.AddInt32(&count, 1)
				fmt.Printf("读第%d次\n",atomic.LoadInt32(&count))
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))

}

func main() {
	testRWlock()
}
