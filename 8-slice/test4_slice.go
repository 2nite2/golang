package main

import "fmt"

func main4() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)
	//len = 4,cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1)
	fmt.Printf("len = %d,cap = %d,slice = %v\n", len(numbers), cap(numbers), numbers)
}
