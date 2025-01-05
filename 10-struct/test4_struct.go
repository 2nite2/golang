package myPackage

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}
type Cat struct {
	Color string
}

func (cat Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}
func (cat Cat) GetColor() string {
	return cat.Color
}
func (cat Cat) GetType() string {
	return "Cat"
}
func Main4() {

}
