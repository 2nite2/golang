package _1_reflect

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTagStr(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagString := t.Field(i).Tag.Get("info")
		fmt.Println("info: ", tagString)
	}
}
func Test4StructTag() {
	var re resume
	findTagStr(&re)
}
