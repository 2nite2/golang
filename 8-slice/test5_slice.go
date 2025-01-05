package main

import "fmt"

func main5() {
	numbers := []int{1, 2, 3, 4, 5}
	s1 := numbers[0:3]
	fmt.Println(s1)
	s1[0] = 100
	// 修改拷贝值时候，原值会一起变化
	fmt.Println(s1)
	s2 := make([]int, 3)
	copy(s2, s1)
	fmt.Println(s2)
}
