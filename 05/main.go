package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)


func main() {
	duration := 5

	ch := make(chan int)

	defer close(ch)

	if len(os.Args) > 1 {
		duration, _ = strconv.Atoi(os.Args[1])
	}

	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	t := time.NewTimer(time.Duration(duration))

	for {
		select {
		case x := <-ch:
			fmt.Println(x)
		case <-t.C:
			return
		}
	}
}
