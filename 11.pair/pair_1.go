package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开指定目录的文件 指定可读可写
	// tty: pair<type:*os.File,value:"/dev/tty">
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if tty != nil {
		fmt.Println("open file error", err)
		return
	}

	// r: pair <type:*io.Reader,value:"/dev/tty">
	var _ io.Reader
	_ = tty
}
