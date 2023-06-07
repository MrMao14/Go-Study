package main

import (
	// l1和l2起别名
	l1 "Go-Study/5.init/lib1"
	l2 "Go-Study/5.init/lib2"
	// 只执行init方法 不使用包内部的方法
	_ "fmt"
	// 导入全部方法 并且可以直接调用方法执行(不建议)
	//. "fmt"
)

func main() {
	l1.Lib1Test()
	l2.Lib2Test()
	/*
		lib1 init()
		lib2 init()
		lib1 test()
		lib2 test()
	*/
}
