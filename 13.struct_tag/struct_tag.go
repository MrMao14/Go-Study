package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"姓名"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagStr := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("tagStr:", tagStr, " ,tagDoc:", tagDoc)
	}
}

func main() {

	var re resume
	// 这里需要传递指针
	findTag(&re)

}
