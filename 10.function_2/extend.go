package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (h *Human) Name() string {
	return h.name
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) Sex() string {
	return h.sex
}

func (h *Human) SetSex(sex string) {
	h.sex = sex
}

func (h *Human) Eat() {
	fmt.Println("Human eat")
}

type SuperMan struct {
	// 继承父类的属性
	Human
	level int
}

func (s *SuperMan) Eat() {
	fmt.Println("SuperMan eat")
}

func main() {
	h := Human{
		name: "zs",
		sex:  "男",
	}

	fmt.Println(h)

	h.Eat()

	s := SuperMan{
		Human: Human{
			name: "zs",
			sex:  "男",
		},
		level: 0,
	}

	fmt.Println(s)

	// 子类重写方法
	s.Eat()
	// 父类方法
	s.SetSex("ss")

}
