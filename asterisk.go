package main

import "fmt"

type Address struct {
	Kota, Provinsi, Wilayah string
}


func main(){
	var address1 Address = Address{"Bogor", "jawa barat", "indonesia"}
var	address2 *Address = &address1//copy value

	address2.Kota = "bandung"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) //berubah menjadi bogor

	// address2 = &Address{"ciawi", "zimbabwe", "tajur"}
	*address2 = Address{"ciawi", "zimbabwe", "tajur"}
	fmt.Println(address1)
	fmt.Println(address2)
}
	
