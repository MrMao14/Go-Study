package lib2

import "fmt"

func init() {
	fmt.Println("lib2 init()")
}

// 当前lib2提供的接口 必须大写(public)
func Lib2Test() {
	fmt.Println("lib2 test()")
}
