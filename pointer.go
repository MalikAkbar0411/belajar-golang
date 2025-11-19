package main

import"fmt"

type Address struct {
	Kota, Provinsi, Wilayah string
}

func main(){
	// address1 := Address{"Bogor", "jawa barat", "indonesia"}
	// address2 := address1 //copy value

	// address2.Kota = "bandung"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) //berubah menjadi bogor

var address1 Address = Address{"Bogor", "jawa barat", "indonesia"}
var	address2 *Address = &address1//copy value

	address2.Kota = "bandung"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) //berubah menjadi bogor
}
	
