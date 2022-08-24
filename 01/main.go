package main

import "fmt"

type (
	Human struct {
		Name string
	}

	Action struct {
		Human
		Name string
	}
)

func (h *Human) Walk() {
	fmt.Println("walk")
}

func main() {
	h := Human{}
	a := Action{Human: h}

	a.Walk()
}
