package main

import "fmt"

func printMap(cityMap map[string]string) {
	fmt.Println("--------------------------")
	for key, value := range cityMap {
		fmt.Println("key = ", key, ",value = ", value)
	}
}
func main2() {
	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "Washington"

	printMap(cityMap)
	delete(cityMap, "USA")
	printMap(cityMap)
}
