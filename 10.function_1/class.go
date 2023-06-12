package main

import "fmt"

type Hero struct {
	Name      string
	Ad, Level int
}

func (h Hero) showName() {
	fmt.Println("Name = ", h.Name)
}

func (h Hero) GetName() string {
	return h.Name
}

func (h Hero) SetName(newName string) {
	h.Name = newName
}

func main() {
	hero := Hero{
		Name:  "zs",
		Ad:    0,
		Level: 1,
	}
	hero.showName()

	// 这里还是值传递
	hero.SetName("lll")

	hero.showName()

	// 这里才是引用传递
	hero.SetName2("sss")

	hero.showName2()

}

func (h *Hero) showName2() {
	fmt.Println("Name = ", h.Name)
}

func (h *Hero) GetName2() string {
	return h.Name
}

func (h *Hero) SetName2(newName string) {
	h.Name = newName
}
