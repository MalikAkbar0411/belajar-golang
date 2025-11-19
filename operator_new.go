package main

import "fmt"

type Address struct {
	Kota, Provinsi, Wilayah string
}


func main ()  {
	var alamat1 *Address = &Address{}
	var alamat2 *Address = alamat1

	alamat2.Wilayah = "zimbabwe"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}