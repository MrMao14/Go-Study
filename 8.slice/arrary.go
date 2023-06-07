package main

import "fmt"

func main() {
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	// range遍历
	for index, value := range myArray2 {
		fmt.Println("index = ", index, ",value = ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
}

// 只能传长度4的数组 并且为值传递!!!
func printArray1(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, ",value = ", value)
	}
}
