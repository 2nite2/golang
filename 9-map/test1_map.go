package main

import "fmt"

func main1() {
	// 1、定义map，分配空间
	var myMap1 map[string]string
	myMap1 = make(map[string]string, 10)
	myMap1["1"] = "java"
	fmt.Println(myMap1["1"])

	// 2、直接分配空间
	myMap2 := make(map[string]string, 10)
	fmt.Println(myMap2)

	//3、直接分配空间，并赋值
	myMap3 := map[string]string{
		"1": "go",
		"2": "java",
		"3": "python",
	}
	fmt.Println(myMap3)

}
