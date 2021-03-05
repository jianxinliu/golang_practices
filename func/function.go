package main

// import . "fmt" 静态导入

import (
	. "fmt"
)

// 函数不支持嵌套、重载和默认参数

// ========================================== 一般函数，符合人的思维顺序 ================================ //
func hello(name string) string {
	return "Hello " + name
}

// ========================================== 返回值命名 ========================================== //
func hello1(name string) (word string) {
	word = "Hello " + name
	return 
}


// ========================================== 多返回值 ========================================== //
func introduction(name ,address string) (word string,addr string){
	word = "Hello " + name
	addr = address
	return 
}

func generater() (int,int){
	return 1,4
}

// ========================================== 函数作为参数 ========================================== //
func test(fn func(x,y int) int,a int,b int) int{
	return fn(a,b)
}

func add(x,y int) int {
	return x + y
}


// 三元运算符的简单替代， 调用方需要自己将返回值转至需要的类型
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// ========================================== 变长参数，传入的是一个 slice ============================== //
func join(split string, str ...string) (ret string) {
	ret = ""
	for i,v := range str {
		ret += If(i != len(str) - 1, v + split, v).(string)
	}
	return 
}



func main(){
	Println(hello("Jack"))
	Println(hello1("Robin"))
	Println(introduction("Pony","China"))
	
	// 支持将函数作为参数传入，函数是第一类对象
	ret := test(add,3,4)
	Println(ret)

	slice := []string{"a","c","d","e","f","g","h"}
	Println(join("-","a","c","d","e","f","g","h"))
	Println(join("=", slice...)) // 直接使用 slice 作为参数，需要手动展开
	
	Println("使用多返回值作为函数的参数：",add(generater()))


	// ========================================== 匿名函数 =============================================== 
	fns := [](func(x int) int){
		func(x int) int {return x},
		func(x int) int {return x*x},
		func(x int) int {return (x + 1)}}

	for _,fn := range fns {
		Println(fn(3))
	}
}
