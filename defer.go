package main

import "fmt"

func logging()  {
	fmt.Println("selesai")
}

func runAplication()  {
	defer logging()
	fmt.Println("run aplikasi")
	
}

func main () {
	runAplication()
}