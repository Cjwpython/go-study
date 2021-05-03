package main

import "fmt"

func main() {
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	slice = arr[2:5]
	fmt.Println(slice)
	fmt.Printf("len:%d,cap:%d \n", len(slice), cap(slice))

	slice = slice[0:1]
	fmt.Println(slice)
	fmt.Printf("len:%d,cap:%d", len(slice), cap(slice))

}
