package _1_reflect

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg), ",value : ", reflect.ValueOf(arg))
}
func MainReflect1() {
	var num float64 = 1.2345
	reflectNum(num)

}
