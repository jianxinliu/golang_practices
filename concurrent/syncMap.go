package main

import (
	"fmt"
	"sync"
)

func main() {
	// 直接声明，不用初始化，不用 make
	var syncMap sync.Map

	// 键值 都是 interface{} 类型
	syncMap.Store("a", 1)
	syncMap.Store("b", 3)
	syncMap.Store(3, 5)

	fmt.Println(syncMap.Load("a"))
	fmt.Println(syncMap.Load(3))
	fmt.Println(syncMap.Load(4))

	syncMap.Delete(3)

	syncMap.Range(func (k, v interface{}) bool {
		fmt.Println(k,v)
		return true
	})
}