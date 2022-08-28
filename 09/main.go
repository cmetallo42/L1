package main

import "fmt"

func main() {
	arr := []int{2, 4, 5, 1, 0, -3, -5, 7, 8, 1, 22, 15, 14, 99, 11}
	ch1, ch2 := make(chan int, len(arr)), make(chan int, len(arr))

	for _, x := range arr {
		ch1 <- x
		ch2 <- int(<-ch1) * 2
		fmt.Println(<-ch2)
	}
}
