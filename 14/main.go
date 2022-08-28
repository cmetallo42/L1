package main

import (
	"fmt"
	"reflect"
)

type types interface {
	i() int
	s() string
	b() bool
	ch() chan any
}

func main() {
	fmt.Println(reflect.TypeOf(types.i))
	fmt.Println(reflect.TypeOf(types.s))
	fmt.Println(reflect.TypeOf(types.b))
	fmt.Println(reflect.TypeOf(types.ch))

}
