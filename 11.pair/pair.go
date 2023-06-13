package main

import "fmt"

func main() {
	var a string

	// pair<static type:string,value:"sss">
	a = "sss"

	// pair<type:string , value:"sss">
	var allType any

	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
