package main

import "fmt"

// fool1函数名称 a,b两个入参 int返回值
func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 匿名返回多个返回值
func foo2(a string, b int) (string, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return a, b
}

// 返回多个返回值 有形参
func foo3(a string, b int) (r1 string, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的返回值变量赋值
	r1 = "cccc"
	r2 = 1000
	return
}

// 返回值类型相同可以统一类型
func foo4(a string, b int) (r1, r2 int) {
	// r1 r2 也属于形参 默认值0
	return
}

func main() {
	c := foo1("abc", 20)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("hh", 99)
	fmt.Println("ret1:", ret1, ",ret2:", ret2)

}
