package main

import "fmt"

// const定义枚举类型
const (
	/*
		可以在const()添加一个关键字 iota 每一行的iota都会累加 第一行默认为0
	*/
	BEIJING  = 10 * iota // iota =0
	SHANGHAI             // iota = 1
	SUZHOU               // iota = 2
)

const (
	a, b = iota + 1, iota + 2 // iota=0 a=iota+1 b=iota+2 a=1 b=2
	c, d                      // iota=1 c=iota+1 b=iota+2 c=2 d=3
	e, f                      // iota=2 e=iota+1 f=iota+2 e=3 f=4

	g, h = iota * 2, iota * 3 // iota=3 g=iota*2 h=iota*3 g=6 h=9
	i, k                      // iota=4 g=iota*2 h=iota*3 g=8 h=12
)

func main() {
	// 常量(只读)
	const length int = 10
	fmt.Println("length = ", length)

	// 常量不允许修改

	fmt.Println("BEIJING:", BEIJING)
	fmt.Println("SHANGHAI:", SHANGHAI)
	fmt.Println("SUZHOU:", SUZHOU)

	// iota只能配合const()使用 iota只在const中进行累加
	//var a int = iota
}
