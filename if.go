package main 

import "fmt"

func main() {
	
	name := "sanusi"

	if name == "abay"{
		fmt.Println("ini abay")	
	}else if name == "malik"{
		fmt.Println("abayaba")
	}else if name == "akbar"{
		fmt.Println("samusi")
	}else {
		fmt.Println("ini bukan abay")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("nama terlalu panjang")
	}else {
		fmt.Println("nama sudah pas ")
	}
}