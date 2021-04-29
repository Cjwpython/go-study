package main

import "fmt"

func main() {
	i := 0
HEAR:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto HEAR
}
