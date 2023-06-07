package main

import "fmt"

func main() {
	// 定义一个切片 长度为3  容量为5
	var numbers = make([]int, 3, 5)

	// len = 3, len = 5, slice = [0 0 0]
	fmt.Printf("len = %d, len = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers中追加一个元素
	numbers = append(numbers, 1)

	// len = 4, len = 5, slice = [0 0 0 1]
	fmt.Printf("len = %d, len = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	/*
		slice会进行自动的动态扩容
		在总数为100以下的时候 会根据cap长度双倍扩容
		超过100之后 1.5倍扩容
	*/

	// len =3 cap = 3
	s := []int{1, 2, 3}

	// 截取数组 最终为[1,2]
	s1 := s[0:2]

	fmt.Println(s1)

	s[0] = 100

	// [100 2 3]
	fmt.Println(s)
	// [100 2]  说明是浅拷贝
	fmt.Println(s1)

	// copy可以将底层数组一起拷贝
	s2 := make([]int, 3)

	copy(s2, s)
	fmt.Println(s2)

}
