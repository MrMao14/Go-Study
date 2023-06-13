package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title: "名称",
		Year:  2023,
		Price: 15,
		Actors: []string{
			"张三",
			"李四",
		},
	}

	// 编码的过程 结构体-->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// json -->结构体
	myMovie := Movie{}

	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error")
		return
	}

	fmt.Printf("%v", myMovie)
}
