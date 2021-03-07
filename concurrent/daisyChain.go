package main

import (
	"fmt"
	"time"
)

// ref https://www.youtube.com/watch?v=f6kdp27TYZs&ab_channel=GoogleDevelopers
// Rob Pike

func f(left, right chan int) {
	left <- 1 + <- right
}

func main() {
	const n = 100000
	leftmost := make(chan int)

	right := leftmost
	left := leftmost

	// build chain
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	// fire
	go func(c chan int) {c <- 1}(right)
	fmt.Println("start...", time.Now())
	fmt.Println(<-leftmost, time.Now())
}