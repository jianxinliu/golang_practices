package main

import (
	"fmt"
)

func change(arr []int) {
	arr[1] = 0
	fmt.Println("inner change: ",arr)
}

func add(arr []int, v int) {
	arr = append(arr, v)
	fmt.Println("inner change: ",arr)
}

func add2(arr *[]int, v int) {
	*arr = append(*arr, v)
	fmt.Println("inner change: ",*arr)
}

func main() {
	// slice
	// a reference to arrays. it dont store any data ,it just describes a section of an underlying array.
	// slice is an array without length limit
	arr := []int{1,2,3,4,5}

	change(arr)

	fmt.Println("after change: ", arr)

	add(arr, 9)

	fmt.Println("after change: ", arr)

	add2(&arr, 9)

	fmt.Println("after change: ", arr)

	fmt.Println("-----------------")

	fmt.Println(arr[1:3])


	// array with fixed length
	var array [3]int

	array[1] = 9

	fmt.Println(array)
}