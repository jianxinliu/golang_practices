package main

import "fmt"
import "sync"

var ch chan  int
var wg *sync.WaitGroup

func elegance(){
	<- ch
	wg.Done()
	fmt.Println("ch get:", 1)
}

func main() {
	wg = &sync.WaitGroup{}
	ch = make(chan int, 3)
	fmt.Println("start.........")
	for i := 0; i < 10; i++ {
		ch <- 1
		wg.Add(1)
		fmt.Println("ch send:", 0)
		go elegance()
	}
	fmt.Println("waiting...........")
	wg.Wait()
	close(ch)
	fmt.Println("done")
}
