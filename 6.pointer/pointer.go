package main

import "fmt"

func main() {
	var a = 1
	// 传入a的地址值
	changeValue(&a)

	fmt.Println("a = ", a)
}

// 将a的地址值存入p中
func changeValue(p *int) {
	// 使用p中存储的地址值找到a所在的栈 找到对应值修改
	*p = 10
}

// 交换两个值
func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}
