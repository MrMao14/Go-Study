package main

import "fmt"

// Animal 本质是一个指针
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

// Cat 定义一个具体的类
type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "Cat"
}

func showAnimal(animal Animal) {
	// 多态
	animal.Sleep()
}

func main() {
	var animal Animal
	animal = &Cat{color: "green"}

	animal.Sleep()
}
