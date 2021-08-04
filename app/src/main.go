package main

import "fmt"

func main() {
	fmt.Println("hello, world")

	x := 3
	y := 6
	result := addition(x, y)
	fmt.Printf("%d+%d=%d", x, y, result)
}

// addition 足し算する
func addition(x int, y int) int {
	return x + y
}
