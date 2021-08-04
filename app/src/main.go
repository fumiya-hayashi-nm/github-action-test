package main

import "fmt"

func main() {
	fmt.Println("hello, world")

	x := 3
	y := 6
	result := plus(x, y)
	fmt.Printf("%d+%d=%d", x, y, result)
}

// plus 足し算する
func plus(x int, y int) int {
	return x + y
}
