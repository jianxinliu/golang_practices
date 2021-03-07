package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // 如果可以写入 c 通道，则计算
			x, y = y, y + x
		case <- quit: // 如果收到介绍消息，则结束
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// 无缓冲通道
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// 循环读 c 通道，是 fibonacci 计算的动力，只有读 ready 写才能 ready, 才能被 select
			fmt.Println(<- c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}