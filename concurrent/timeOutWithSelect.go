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
			case <- timeout: // 如果 ch 没有 ready ,而 timeout 计时到达，则认为 ch 没有在规定时间内 ready ,则认为超时
				fmt.Println("time out")
		}
	}
}

