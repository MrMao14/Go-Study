package main

import "fmt"

func main() {
	// 声明slice1是一个切片 并初始化 默认值1,2,3 长度3
	//slice1 := []int{1, 2, 3}

	// 声明slice1是一个切片 但是没有分配空间
	//var slice1 []int
	// make来开辟空间 默认0
	//slice1 = make([]int, 3)

	// 声明slice是一个切片 并分配空间 初始化0
	var slice1 []int = make([]int, 3)
	//var slice1 := make([]int, 3)

	fmt.Printf("myArray2 types = %v\nlen = %d", slice1, len(slice1))

	// 判断一个slice是否为0
	if slice1 == nil {
		fmt.Println("slice1 为 null")
	} else {
		fmt.Println("slice1 不为 null")
	}
}
