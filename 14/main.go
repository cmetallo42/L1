package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := any(int(1))
	s := any("hello world")
	b := any(true)
	ca := any(make(chan any))

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(ca))

}
