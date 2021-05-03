package main

import (
	"fmt"
	"sort"
)

func testMap() {
	a := make(map[int]int, 5)
	a[1] = 5
	a[2] = 3
	a[3] = 2
	a[4] = 4
	a[5] = 1

	for k, v := range a {
		fmt.Println("k:", k, "value", v)
	}
	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}

}

func main() {
	testMap()
}
