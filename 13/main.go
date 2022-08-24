package main

import "fmt"

func swap(x int, y int) {
	x = x + y
	y = x - y
	x = x - y
	fmt.Println(x, y)
}

func swap2(x int, y int) {
	x, y = y, x
	fmt.Println(x, y)
}

func main() {
	x, y := 5, 7

	swap(x, y)
	swap2(x,y)
}
