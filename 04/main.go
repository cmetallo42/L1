package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	number := 4

	if len(os.Args) > 1 {
		number, _ = strconv.Atoi(os.Args[1])
	}

	important := make(chan any, number)

	for i := 0; i < number; i++ {
		go func() {
			for {
				fmt.Println(<-important)
			}
		}()
	}

	go func() {
		for {
			important <- rand.Int()
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals
}
