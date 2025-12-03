package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main(){
	result := helper.SayHello("abay")
	fmt.Println(result)

	fmt.Println(helper.Aplikasi)
	// fmt.Println(helper.version) //tidak bisa di akses
	// fmt.Println(helper.sayGoodBye) tidak bisa di akses karena hurus awalnya kecil
}
