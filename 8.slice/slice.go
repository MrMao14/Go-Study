package main

import "fmt"

func main() {
	// 动态数据 切片slice
	myArray := []int{1, 2, 3, 4, 5}

	fmt.Printf("myArray2 types = %T\n", myArray)

	printArray2(myArray)
}

func printArray2(myArray []int) {
	// _ 表示匿名变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	// 引用传递 可以修改值
	myArray[0] = 100
}
