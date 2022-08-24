package main

import (
	"fmt"
	"time"
)

func worker(i int, result chan<- int) {
	fmt.Println("Starting", i, "*", i, "operation.")
	time.Sleep(1 * time.Second)
	result <- i * i
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	result := make(chan int)
	total := 0

	for i := 0; i < len(nums); i++ {
		go worker(nums[i], result)
	}
	for j := 0; j < len(nums); j++ {
		total += <-result
	}
	fmt.Println("Total:", total)
}
