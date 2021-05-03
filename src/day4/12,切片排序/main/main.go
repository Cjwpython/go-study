package main

import (
	"fmt"
	"sort"
)

func testintSort() {
	var a = [...]int{1, 2, 3, 5, 4}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testStringSort() {
	var a = [...]string{"weqe", "vsdvsd", "a", "b", "c"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func testFloatSort() {
	var a = [...]float64{2.4, 3.4, 5.5, 6.6, 3.3}
	sort.Float64s(a[:])
	fmt.Println(a)
}

func testIntSearch() {
	var a = [...]int{1, 2, 3, 4, 8, 5, 6}
	index := sort.SearchInts(a[:], 1)
	fmt.Println(a[index])
}
func main() {
	//testStringSort()
	//testintSort()
	//testFloatSort()
	testIntSearch()
}
