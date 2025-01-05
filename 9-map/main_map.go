package main

import myPackage "GoLangStudy/10-struct"

func main() {
	hero := myPackage.Hero{Level: 1}
	hero.Show()
	hero.SetName("azhen")
}
