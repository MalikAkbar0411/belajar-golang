package main

import "fmt"

type Customer struct{
	Name string
	Address string
	Age int 
}

func main()  {
	var abay Customer
	fmt.Println(abay)

	abay.Name = "abay "
	abay.Address = "indonesia"
	abay.Age = 17
	
	fmt.Println(abay)

	abay = Customer{
		Name : "abay",
		Address : "indonesai",
		Age : 30,
	}
	fmt.Println(abay)

	budi := Customer{"budi", "indonesia", 30}
	fmt.Println(budi)
}