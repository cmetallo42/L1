package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr1, arr2 := []int{}, []int{}
	m := make(map[int]struct{})
	for i := 0; i < 10; i++ {
		arr1 = append(arr1, rand.Intn(20))
		arr2 = append(arr2, rand.Intn(20))
		m[arr1[i]] = struct{}{}
		m[arr2[i]] = struct{}{}
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	for k := range m {
		fmt.Print(k, " ")
	}
	fmt.Println("")
}
