package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notfounderror struct {
	Message string
}

func (n *notfounderror) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "abay" {
		return &notfounderror{"data not found"}

	}
	return nil
}
func main() {
	err := SaveData("abay", nil)
	if err != nil {
		//terjadinerror
		if validationErr, ok := err.(*validationError); ok{
			fmt.Println("validation error:", validationErr.Error())
		}else if notfounderror, ok := err.(*notfounderror); ok{
			fmt.Println("not found error:", notfounderror.Error())
		}else{
			fmt.Println("unkwon error:", err.Error())
		}
	} else {
		//silses
		fmt.Println("sukses")
	}
}
