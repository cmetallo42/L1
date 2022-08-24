package main

import (
	"fmt"
	"time"
)

func worker(i int, results chan<- int) {
	fmt.Println("Starting")
	time.Sleep(1 * time.Second)
	results <- i * i
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(nums))

	for i := 0; i < len(nums); i++ {
		go worker(nums[i], results)
	}
	for j := 1; j <= len(nums); j++ {
		fmt.Println(<- results)
		fmt.Println("Done")
	}
}
