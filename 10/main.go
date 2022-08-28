package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float32)
	tmpInt := 0
	tmpStr := ""

	for _, i := range arr {
		tmpInt = int(i) / 10
		tmpStr = strconv.Itoa(tmpInt) + "0"
		tmpInt, _ = strconv.Atoi(tmpStr)
		v := m[tmpInt]
		v = append(v, i)
		m[tmpInt] = v
	}
	fmt.Println(m)
}
