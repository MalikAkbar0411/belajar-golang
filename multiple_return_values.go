package main

import "fmt"

func getFullName() (string, string){
	return "abay", "malik"
}

// func main()  {
// 	firstName, lastName := getFullName()
// 	fmt.Println(firstName, lastName)
// }

func main()  {//cuma ngambil awal doang
	firstName, _ := getFullName()
	fmt.Println(firstName)
}