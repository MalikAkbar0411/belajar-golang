package main

import "fmt"

type filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string)  {
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

func spamFilter(name string) string {
	if name =="anjing" {
		return "..."
	}else {
		return name
	}
}

func main()  {
	sayHelloWithFilter("abay", spamFilter)
	 filter := spamFilter
	 sayHelloWithFilter("anjing", filter)
}