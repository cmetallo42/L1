package main

import "fmt"

func main() {
	s1 := "abcdef"
	var s2 string

	for i := len(s1) - 1; i >= 0; i-- {
		s2 += string(s1[i])
	}
	
	fmt.Println(s2)
}
