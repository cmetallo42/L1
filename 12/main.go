package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})

	for _, i := range arr {
		m[i] = struct{}{}
	}
	for k := range m {
		fmt.Print(k, " ")
	}
	fmt.Println("")
}
