package main

import "fmt"

func main() {
	 var names[3]string

	 names [0] = "abay"
	 names [1] = "malik"
	 names [2] = "akbar"

	 fmt.Println(names[0])
	 fmt.Println(names[1])
	 fmt.Println(names[2])

	var values = [...]int{
		90,
		21,
		23,
		23,
		12,
		12,
	}
	fmt.Println(values)
	fmt.Println( len(values))

}