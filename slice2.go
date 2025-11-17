package main

import "fmt"
func main() {

	names := [...] string{"abay", "malik", "akbar" , "sanusi", "huge", "haf"}

	slice1 := names[2:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	var slice4[] string = names[:] //manual
	fmt.Println(slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days) 

		daySlice2 := append(daySlice1, "libur baru")
		daySlice2[0] = "sabtu lama"
		fmt.Println(daySlice2)
		fmt.Println(daySlice1)

		var newSlice []string = make([]string, 2,5)
		newSlice [0] = "abay"
		newSlice [1] = "abay"
		// newSlice [2] = "abay" error, harusnya pake append

		fmt.Println(newSlice)
		fmt.Println(len(newSlice))
		fmt.Println(cap(newSlice))

		newSlice2 := append(newSlice, "abay")
		fmt.Println(newSlice2)
		fmt.Println(len(newSlice2))
		fmt.Println(cap(newSlice2))

		newSlice2[0] = "budi"
		fmt.Println(newSlice2)
		fmt.Println(newSlice)

		fromSlice := days[:]
		toSlice := make([]string, len(fromSlice), cap(fromSlice))
		
		copy(toSlice, fromSlice)

		fmt.Println(fromSlice)
		fmt.Println(toSlice)

		iniArray := [...]int{1, 2, 3, 4, 5}
		iniSlice := []int {1,2,3,4,5}

		fmt.Println(iniArray)
		fmt.Println(iniSlice)

}

