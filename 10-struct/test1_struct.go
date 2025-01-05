package myPackage

import "fmt"

type myInt int

type Book struct {
	title string
	auth  string
}

func changeBook1(book Book) {
	book.title = "java"
}

func changeBook2(book *Book) {
	book.title = "java"
}
func main1() {
	fmt.Println(111)
	var a myInt
	a = 1
	fmt.Printf("Type is : %T\n", a)
	var book1 Book
	book1.title = "go"
	book1.auth = "zhen"
	fmt.Printf("%v\n", book1)
	changeBook1(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)

}
