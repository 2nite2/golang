package main

import "fmt"

func main2() {
	// 1、声明切片，初始化
	slice1 := []int{1, 2, 3}
	// 2、 声明切片，但不分配空间,通过make分配
	var slice2 []int
	slice2 = make([]int, 4)
	slice2[0] = 10
	// 3、声明切片，同时分配空间
	var slice3 []int = make([]int, 3)
	// 4、声明切片，同时分配空间
	slice4 := make([]int, 3)
	fmt.Printf("len = %d,slice = %v\n", len(slice1), slice1)
	fmt.Printf("len = %d,slice = %v\n", len(slice2), slice2)
	fmt.Printf("len = %d,slice = %v\n", len(slice3), slice3)
	fmt.Printf("len = %d,slice = %v\n", len(slice4), slice4)

}
