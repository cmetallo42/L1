package main

import (
	"fmt"

	"main/16"
)

func BinaryFind(arr []int, f int) {
	fmt.Println(arr)
	if len(arr) < 1 {
		fmt.Println("Empty Array")
		return
	}
	if len(arr) < 3 {
		if len(arr) == 1 {
			if arr[0] == f {
				fmt.Println("Founded")
			} else {
				fmt.Println("Not Founded")
			}
		} else {
			if arr[0] == f || arr[1] == f {
				fmt.Println("Founded")
			} else {
				fmt.Println("Not Founded")
			}
		}
		return
	}

	mid := len(arr) / 2
	if arr[mid] > f {
		BinaryFind(arr[:mid], f)
	} else if arr[mid] < f {
		BinaryFind(arr[mid:], f)
	} else {
		fmt.Println("Founded")
	}
}

func main() {
	arr := []int{2, 4, 5, 1, 0, -3, -5, 7, 8, 1, 22, 15, 14, 99, 87}
	sortedArr := sixteen.QuickSort(arr)
	BinaryFind(sortedArr, 7)
}
