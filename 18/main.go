package main

import (
	"strconv"
	"sync"
)

type Incrementor struct {
	mu     sync.Mutex
	number int
}

func (i *Incrementor) Add() {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.number++
}

func (i *Incrementor) String() string {
	i.mu.Lock()
	defer i.mu.Unlock()
	
	return strconv.Itoa(i.number)
}

func main() {
	inc := &Incrementor{}

	wg := sync.WaitGroup{}

	wg.Add(4)

	for i := 0; i < 4; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				inc.Add()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	println(inc.String())
}
