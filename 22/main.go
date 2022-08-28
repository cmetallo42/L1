package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("1000000000000000000000000000000000000000000000000", 10)
	b, _ := new(big.Int).SetString("500000000000000000000000000000000000000000000000", 10)
	z := new(big.Int)

	fmt.Println(z.Mul(a, b))
	fmt.Println(z.Div(a, b))
	fmt.Println(z.Add(a, b))
	fmt.Println(z.Sub(a, b))
}
