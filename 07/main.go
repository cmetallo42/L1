package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func main() {
	workers := 4
	quantity := 25

	var m sync.Map
	var wg sync.WaitGroup

	if len(os.Args) > 1 {
		workers, _ = strconv.Atoi(os.Args[1])
	}

	if len(os.Args) > 2 {
		quantity, _ = strconv.Atoi(os.Args[2])
	}

	wg.Add(workers * quantity)

	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < quantity; j++ {
				m.Store(rand.Int(), rand.Int())
				wg.Done()
			}
		}()
	}
	wg.Wait()

	total := 0
	m.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, "\tval:", v)
		total++
		return true
	})

	fmt.Println("Total:", total)
}
