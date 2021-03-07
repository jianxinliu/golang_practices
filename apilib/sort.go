package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (bl ByLength) Len() int {
	return len(bl)
}

func (bl ByLength) Swap(i, j int) {
	bl[i], bl[j] = bl[j], bl[i]
}

// 如何比较
func (bl ByLength) Less(i, j int) bool {
	return len(bl[i]) < len(bl[j])
}


func main() {
	strs := []string{"1", "5", "2", "3", "9"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{3,1,5,7,3,5}
	sort.Ints(ints)
	fmt.Println(ints)

	fmt.Println(sort.IntsAreSorted(ints))

	// 自定义排序
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}