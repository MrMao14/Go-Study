package main

import "fmt"

// 声明一种行的数据类型 myInt是int的一个别名
type myInt int

// 定义一个结构体
type Book struct {
	title string
	auth  string
}

func main() {

	var book1 Book
	book1.title = "golang"
	book1.auth = "mrmao"

	fmt.Printf("%v\n", book1)

	changeBook(book1)

	// 值不会改变 值传递
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	// 传递指针 值改变
	fmt.Printf("%v\n", book1)

}

func changeBook(book Book) {
	// 传递一个book的副本
	book.auth = "666"
}

func changeBook2(book *Book) {
	book.auth = "777"
}
