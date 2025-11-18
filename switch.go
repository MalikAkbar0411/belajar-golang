package main

import "fmt"

func main ()  {

	name  := "waawawa"
	switch name {
	case "abay":
		fmt.Println("hello abay")
	case "malik":
		fmt.Println("hello malik")
	default:
	fmt.Println("wakawaw")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama pas")

	}

	name = "abay"
	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("nama terlalu panjang")
	case lenght > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("sudah benar")

	}
}