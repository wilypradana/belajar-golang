package main

import "fmt"

func main() {
	// 	// fmt.Println(len(slice3))
	// fmt.Println(cap(slice1))
	names := [...]string{"eko","kurniawan","khannedy","joko","budi","nugraha"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)
	
	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"senin","selasa","rabu","kamis","jumat","sabtu","minggu"}
	daySlice1 := days[5:] // sabtu,minggu
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days)
	
	daySlice2 := append(daySlice1, "Libur baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2) //  [...]string{"senin","selasa","rabu","kamis","jumat","sabtu","minggu","Libur baru"}
	fmt.Println(days)

	newSlice  :=  make([]string, 2, 5)

	newSlice[0] = "Eko"
	newSlice[1] = "kurniawan"
	newSlice2 := append(newSlice, "khannedy")
	newSlice3 := append(newSlice, "budi")
	newSlice4 := append(newSlice, "nugraha")

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(newSlice2)
	fmt.Println(newSlice3)
	fmt.Println(newSlice4)




}