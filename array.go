package main

import "fmt"

func main() { 
	var names [4]string


names[0] = "abay"
names[1] = "malik"
names[2] = "akbar"
names[3] = "sanusi"

fmt.Println(names[0])
fmt.Println(names[1])
fmt.Println(names[2])
fmt.Println(names[3])

var values = [...]int{
	90,
	80,
	70,
	100,
}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}