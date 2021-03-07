package main

import (
	"fmt"
	"strings"
)

// 使用 go 实现数据处理的函数 (string 版)

func Index(arr []string, t string) int {
	for i, v := range arr {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(arr []string, t string) bool {
	return Index(arr, t) > -1
}

func Any(arr []string, fn func(string) bool) bool {
	for _, v := range arr {
		if fn(v) {
			return true
		}
	}
	return false
}

func All(arr []string, fn func(string) bool) bool {
	for _, v := range arr {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Filter(arr []string, fn func(string) bool) []string {
	ret := make([]string, 0)
	for _, v := range arr {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map(arr []string, fn func(string) string) []string {
	ret := make([]string, len(arr))
	for i, v := range arr {
			ret[i] = fn(v)
	}
	return ret
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
			return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
			return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
			return strings.Contains(v, "e")
	}))

	// 上面的例子都是用的匿名函数，但是你也可以使用类型正确的命名函数
	fmt.Println(Map(strs, strings.ToUpper))
}