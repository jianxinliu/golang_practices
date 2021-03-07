package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	ch := make(chan int)
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			rd := rand.Intn(10)
			if rd <= 5 {
				ch <- rd
			} else {
				// 0.5 的概率发生延迟
				time.Sleep(time.Second * 3)
			}
		}
	}()

	for {
		// 返回一个 channel， 该 channel 在指定时间后会传入一个值，在这之前一直阻塞
		timeout := time.After(2 * time.Second)

		// 比哪个 channel 的阻塞时间短
		select {
			case s := <-ch:
				fmt.Println(s)
			case <- timeout: // 如果 ch 没有 ready ,而 timeout 计时到达，则认为 ch 没有在规定时间内 ready ,则认为超时
				fmt.Println("time out")
		}
	}
}

