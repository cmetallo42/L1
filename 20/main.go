package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "snow dog sun"
	s2 := strings.Split(s1, " ")
	var s3 string
	for i := len(s2)-1; i >= 0; i-- {
		s3 += s2[i] + " "
	}
	s4 := strings.Trim(s3, " ")
	
	fmt.Println(s4)
}
