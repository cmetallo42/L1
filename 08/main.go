package main

import (
	"fmt"
	"strconv"
)

func main() {
	number64 := int64(123)
	var newStrBase string
	k := 4

	strBase := strconv.FormatInt(number64, 2)
	if k > len(strBase)-1 || k < 0{
		fmt.Println("Index must be in range")
		return
	}

	fmt.Println(strBase)

	newStrBase += strBase[:k]

	if strBase[k] == '0' {
		newStrBase += "1"
	} else {
		newStrBase += "0"
	}

	newStrBase += strBase[k+1:]

	fmt.Println(newStrBase)

	if newNumber64, err := strconv.ParseInt(newStrBase, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newNumber64)
	}
}
