package main

import "fmt"

func main1() {
	MyArray1 := [4]int{1, 2, 3, 4}
	for index, value := range MyArray1 {
		fmt.Println("index = ", index, ",value = ", value)
	}
	MyArray2 := [10]int{4, 5, 6, 7}
	fmt.Printf("MyArray1 Type is %T\n", MyArray1)
	fmt.Printf("MyArray2 Type is %T\n", MyArray2)
}
