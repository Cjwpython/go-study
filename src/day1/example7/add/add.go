package add

func Add(a int, b int, c chan int) {
	var sum int
	sum = a + b
	c <- sum

}
