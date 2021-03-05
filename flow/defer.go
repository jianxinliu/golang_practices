package main

import "fmt"

func main() {
	// defer 延迟处理，多个 defer 语句的执行顺序是 FILO, 同栈结构
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	// return 会做的几件事：
	// 1. 给返回值赋值
	// 2. 调用 defer 表达式
	// 3. 返回给调用函数

	for i := 0; i < 5; i++ {
		defer func() {
			// 延迟处理，最后取得 i 为 5
			fmt.Println(i)
		}()
	}

	for i := 0; i < 5; i++ {
		defer func(idx int) {
			fmt.Println(idx)
		}(i) // 传入的值会被即可求值保存
	}

}
