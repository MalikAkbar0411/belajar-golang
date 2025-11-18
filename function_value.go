package main

import "fmt"

func getGoodBye(name string)string  {
	return "GOOD bye" + name
}

func main()  {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh(" abay"))
	fmt.Println(misal(" malik"))
}