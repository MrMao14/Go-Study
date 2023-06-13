package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) Call() {
	fmt.Println("user is call")
	fmt.Printf("%v\n", u)
}

func main() {
	user := User{
		Id:   1,
		Name: "zs",
		Age:  18,
	}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input any) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType:", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue:", inputValue)

	// 通过type获取里面的字段
	// 1.获取any的reflect,type 通过Type得到NumField
	// 2.得到每个field数据类型
	// 3.通过field有Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s : %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type获取里面的方法进行调用
	// Go 1.17之后无效
	for i := 0; i < inputType.NumField(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s : %v\n", m.Name, m.Type)
	}
}
