package lib1

import "fmt"

func init() {
	fmt.Println("lib1 init()")
}

// 当前lib1提供的接口 必须大写(public)
func Lib1Test() {
	fmt.Println("lib1 test()")
}
