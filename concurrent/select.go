package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	go func() {
		for {
			fmt.Println("get:", <-ch1)
		}
	}()

	for {
		// 向 channel 随机发送
		select {
			case ch1 <- 0:
			case ch1 <- 1:
		}
		time.Sleep(time.Second * 2)
	}
}
