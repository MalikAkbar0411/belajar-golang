package  main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	

	// person["name"] = "abay"
	// person["address"] = "bogor"

	person:=map[string]string{
	"name" : "abay",
	"address" : "bogor",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "book golang"
	book["author"] = "abay"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")
	
	fmt.Println(book)
 }