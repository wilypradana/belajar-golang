package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string)string)  {
// 	fmt.Println("Hello", filter(name))
// }

// // ini akan terafiliasi dengan filter func di paramater sayhelloWithFilter
// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	}else{
// 		return name
// 	}
// }

// func main() {
// 	udahDifilter := spamFilter
// 	sayHelloWithFilter("Anjing", udahDifilter)
// }


func sayHelloWithFilter(name string, age int, filterName func(string) string, filterAge func(int) string)  {
	fmt.Println("Hello", filterName(name),filterAge(age))
}

func filterNames(name string) string {
	if name == "Anjing" {
		return "..."		
	}else{
		return name
	}
}

func filterAges(age int) string {
	if age < 18 {
		return "maaf kamu belum cukup umur"		
	}else{
		return "selamat kamu sudah cukup umur"
	}
}

func main()  {
	filterName :=  filterNames
	filterAge := filterAges
	sayHelloWithFilter("Joko", 18,filterName,filterAge)
}