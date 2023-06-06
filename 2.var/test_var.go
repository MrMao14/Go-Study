package main

/*
变量的声明方式
*/
import "fmt"

// 声明全局变量
var gA int = 100
var gB = 200

// := 只能使用在函数体内部使用

func main() {
	// 方法1 声明一个变量 默认0
	var a int
	fmt.Println("a = ", a)
	// 输出数据类型
	fmt.Printf("type of a = %T\n", a)

	// 方法2 声明一个变量 并初始化
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abc"
	fmt.Printf("bb = %s type of bb = %T\n", bb, bb)

	// 方法3 初始化时可以省去数据类型 自动匹配数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abc"
	fmt.Printf("cc = %s type of cc = %T\n", cc, cc)

	// 方法4 (常用) 省去var关键字 直接匹配
	e := 100
	fmt.Println("e = ", e)

	// ============
	fmt.Println("gA = ", gA, ", gB = ", gB)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)

	var kk, ll = 100, "ssss"
	fmt.Println("kk = ", kk, ",ll = ", ll)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ",jj = ", jj)
}
