package main

import "fmt"

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
func main() {
	a := 100
	b := 200
	swap(&a, &b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	a = 100
	b = 200
	a, b = b, a
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
