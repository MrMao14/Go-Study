package main

import "fmt"

func main() {
	/*
		defer关键字 相当于finally最后执行
		相当于压栈的方式
		2相比较于1会先执行
	*/
	defer fmt.Println("defer::hello 1")
	defer fmt.Println("defer::hello 2")

	fmt.Println("main::hello  1")
	fmt.Println("main::hello  2")

	/*
		return会优先于defer调用
		并且会先执行内部所有方法 再到main函数的defer
	*/
	returnAndDefer()

	/*
		main::hello  1
		main::hello  2
		return...
		defer..
		defer::hello 2
		defer::hello 1
	*/
}

func deferFunc() int {
	fmt.Println("defer..")
	return 0
}

func returnFunc() int {
	fmt.Println("return...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()

	return returnFunc()
}
