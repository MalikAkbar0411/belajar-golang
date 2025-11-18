package main

import "fmt"

func main(){

	for i := 0; i < 11; i++{
		 if i == 2 {
			break
		 }
		 fmt.Println("perulangan ke", i)
	}
}