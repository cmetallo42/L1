package main

import (
	"fmt"
	"main/25"
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

	go func(){
		for {
			fmt.Println(<-ch)
		}
	}()


	twentyfive.Sleep2(time.Duration(duration))
}
