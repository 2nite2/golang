package myPackage

import "fmt"

func main() {
	//main2()
	hero := Hero{name: "huang", id: 2, Level: 100}
	hero.Show()
	fmt.Println(hero.id)
}
