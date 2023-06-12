package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	cityMap["china"] = "beijing"
	cityMap["japan"] = "tokyo"
	cityMap["usa"] = "new york"

	// 遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}

	// 删除
	delete(cityMap, "china")

	// 修改
	cityMap["usa"] = "DC"

	fmt.Println("==================")
	printMap(cityMap)

	fmt.Println("==================")
	//传值修改
	changeValue(cityMap)
	printMap(cityMap)
}

func printMap(cityMap map[string]string) {
	// cityMap是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func changeValue(cityMap map[string]string) {
	cityMap["england"] = "london"
}
