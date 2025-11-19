package main

import "fmt"

type Address struct {
	Kota, Provinsi, Wilayah string
}


func ChangeCountryToIndonesia(Address *Address)  {
		Address.Wilayah = "waluh"
}

func main()  {
	var address Address = Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}