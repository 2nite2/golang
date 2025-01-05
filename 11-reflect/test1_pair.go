package _1_reflect

import (
	"fmt"
)

func Main1() {
	var str = "abcd"
	var alltype interface{}
	alltype = str
	value, _ := alltype.(string)
	fmt.Println(value)
}
