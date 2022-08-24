package main

import "fmt"

func Remove(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{2, 4, 5, 1, 0, -3, -5, 7, 8, 1, 22, 15, 14, 99, 87}
	Remove(slice, 2)
	fmt.Println(slice)
}
