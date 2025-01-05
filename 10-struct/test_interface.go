package myPackage

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println(arg)
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value = ", value)
	}
}

type BookTip struct {
	auth string
}

func Main5() {
	book := BookTip{"Golang"}
	myFunc(book)
	myFunc("111")
}
