package main 
import "fmt"

func main() {
	names := [...]string{"abay", "muhammad", "malik", "akbar", "sanusi", "bruno"}

	slice1 := names[3:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	var slice3 []string= names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)


	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daysSlice1 := days[5:] 
	fmt.Println(daysSlice1)

	daysSlice1[0]= "sabtu ngopi"
	daysSlice1[1]= "minggu turu"

	fmt.Println(daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "nongkrong")
	daysSlice2[0]= "sabtu molor"
	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)

	var newslice [] string	= make([]string, 2 , 5)
	newslice[0]= "abay"
	newslice[1]= "abay"
	// newslice[2]= "abay" ??error harusnya menggunakan append
	

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	newslice2 := append(newslice, "abay")
	fmt.Println(newslice2)
	fmt.Println(len(newslice2))
	fmt.Println(cap(newslice2))

	newslice2[0] = "yaba"
	fmt.Println(newslice2)
	fmt.Println(newslice)


	//copy slice
	fromslice := names[:]
	toSlice := make([]string, len(fromslice), cap(fromslice))
	
	copy(toSlice, fromslice)
	fmt.Println(fromslice)
	fmt.Println(toSlice)

	// bedanya array dan slice
	iniArray := [...]int{1,2,3,4,5,6}
	iniSlice := []int{1,2,3,4,5,6,}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
