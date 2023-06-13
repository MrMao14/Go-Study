package main

import "fmt"

// interface{}是万能数据类型
func myFunc1(arg any) {
	// any等同于interface{}
}

func myFunc(arg interface{}) {
	fmt.Println("myFunc....")
	fmt.Println(arg)

	// interface()如何区分数据类型
	value, ok := arg.(string)

	if ok {
		fmt.Println("arg是string")
	} else {
		fmt.Println("arg不是string")
		fmt.Println(value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{auth: "Golang"}
	myFunc(book)

	myFunc(100)

	myFunc("abc")

	myFunc(3.14)
}
