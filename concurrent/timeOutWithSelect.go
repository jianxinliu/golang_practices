package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	ch := make(chan int)
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			rd := rand.Intn(10)
			if rd <= 5 {
				ch <- rd
			} else {
				time.Sleep(time.Second * 3)
			}
		}
	}()

	timeout := time.After(2 * time.Second)

	for {
		select {
			case s := <-ch:
				fmt.Println(s)
			case <- timeout:
				fmt.Println("time out")
		}
	}
}

