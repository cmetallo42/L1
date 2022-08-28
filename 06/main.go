package main

import (
	"fmt"
	"sync"
	"main/25"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	ch := make(chan int)
	flag := true

	go func(){
		for {
			select {
			case <- ch:
				fmt.Println("Goroutine1 has stopped by channel")
				wg.Done()
				return
			default:
				fmt.Println("Goroutine1 working...")
				twentyfive.Sleep2(1)
			}
		}
	}()

	go func(){
		for flag {
			fmt.Println("Goroutine2 working...")
			twentyfive.Sleep2(1)
		}
		fmt.Println("Goroutine2 has stopped by flag")
		wg.Done()
	}()

	twentyfive.Sleep2(5)

	ch <- -1
	flag = false

	close(ch)
	wg.Wait()
}
