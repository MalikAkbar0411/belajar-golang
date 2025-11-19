package main

import "fmt"

func NewMap(name string)map [string] string  {
	if name == ""{
		return nil
	}else {
		return map [string]string{
			"name" : name,
		}
	}
}

func main()  {
	data := NewMap("abay")
	if data == nil {
		fmt.Println("data mao masih kosong")
	}else{
		fmt.Println(data["name"])
	}
} 