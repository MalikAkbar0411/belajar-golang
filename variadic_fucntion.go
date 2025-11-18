package main

import "fmt"

func sumALL(numbers ... int)int{	
	total := 0
	
	for _, number := range numbers{
		total += number
	}

	return total
}
	func main()  {
		fmt.Println(sumALL(10,10,10))
		fmt.Println(sumALL(10,10,10,10))
		fmt.Println(sumALL(10,10,10,10,10))

		numbers := [] int {10, 10, 10 ,10}
		fmt.Println(sumALL(numbers...))

	}
