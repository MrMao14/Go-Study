package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.234556
	reflectNum(num)
}

func reflectNum(arg any) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value:", reflect.ValueOf(arg))
}
