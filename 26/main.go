package main

import (
	"fmt"
	"strings"
)

func CheckUnique(s1 string) bool  {
	s1 = strings.ToLower(s1)

	collection := make(map[rune]struct{})

	for _, r := range s1 {
		_, ok := collection[r]
		if ok {
			return false
		}

		collection[r] = struct{}{}
	}

	return true
}

func main() {
	s1 := "abCdejhl"
	b := CheckUnique(s1)
	fmt.Println(b)
}
