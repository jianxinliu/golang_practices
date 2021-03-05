package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("send", i)
			ch <- i
		}
	}()

	go func() {
		for y := range ch {
			wg.Done()
			fmt.Println("get ", y)
		}
	}()

	wg.Wait()
	close(ch)
}
