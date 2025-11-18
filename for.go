package main

import "fmt"

func main (){
	// counter := 1
	// for counter <= 100{
	// 	fmt.Println("perulangan ke", counter)
	// 	counter++
	// }

	//lebih singkat
	for counter := 1; counter <= 10; counter++{
		fmt.Println("perulangan ke", counter)
	}

	fmt.Println("selesai")

	names := []string{"abay", "malik", "akbar"}
	for i := 0; i <len(names); i++ {
		fmt.Println(names[i])
	}

	//cara singkat

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
 }