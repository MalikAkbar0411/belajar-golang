package main

import "fmt"

func getCompleteName() (firstName , middleName , lastName string)  {
	firstName = "abay"
	middleName = "malik"
	lastName  = "sanusi"
	
	return firstName, middleName, lastName
} 

func main()  {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}