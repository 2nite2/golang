package myPackage

import "fmt"

// 如果类名首字母大写，表示属性为public
type Hero struct {
	name  string
	id    int
	Level int
}

func (h *Hero) GetName() string {
	return h.name
}
func (h *Hero) SetName(newName string) {
	h.name = newName
}
func (h *Hero) Show() {
	fmt.Println("name =", h.name, ", id =", h.id, ", Level =", h.Level)
}

func Main2() {
	hero := Hero{name: "zhen", id: 1, Level: 10}
	hero.Show()
	hero.SetName("huang")
	hero.Show()
}
