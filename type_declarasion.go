package main

import "fmt"

func main() {

	type NoKTP string

	var ktpAbay NoKTP = "1111111111"

	var contoh string = "2222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpAbay)
	fmt.Println(contohKtp)
}