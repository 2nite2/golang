package _1_reflect

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", user)
}

func MainReflect2() {
	user := User{1, "2nite2", 25}
	DoFieldAndMethod(user)
}

func DoFieldAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputType is :", inputType.Name(), ",inputValue is :", inputValue)
	fmt.Println("---------------------------------------------------")
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println("field.Name is :", field.Name, ",field.Type is :", field.Type, ",value is :", value)
	}
	fmt.Println("---------------------------------------------------")

	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Println("method.Name is :", method.Name, ",method.Type is :", method.Type)
	}
}
