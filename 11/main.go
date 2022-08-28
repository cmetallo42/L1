package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	arr1, arr2 := []int{}, []int{}
	m := make(map[int]struct{})
	for i := 0; i < 10; i++ {
		arr1 = append(arr1, r.Intn(20))
		arr2 = append(arr2, r.Intn(20))
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
