package main

import "fmt"

func getHello(name string) string{
	hello := "Hello" + name
	return hello

}


func main(){
	result:= getHello("abay")
	fmt.Println(result)

	fmt.Println(getHello("akbar"))
		fmt.Println(getHello(",alik"))

}		