package _1_reflect

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year""`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func Test6_json() {
	movie := Movie{Title: "你好", Year: 2024, Price: 40, Actors: []string{"zhaoying", "huangsheng"}}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	movie2 := Movie{}
	err = json.Unmarshal(jsonStr, &movie2)
	if err != nil {
		fmt.Println("json Unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", movie2)

}
