package main

import "fmt"

type Customer struct{
	Name,Address string
	Age int
}

func (customer Customer) sayHelloWithFilter(name string, age int, filterName func(string) string, filterAge func(int) string)  {
	fmt.Println("Hello", filterName(name),filterAge(age),age,  "aku", customer.Name, "customer service")
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
		return "maaf kamu belum cukup umur umur kamu"		
	}else{
		return "umur kamu"
	}
}

func main() {
	budi := Customer{"Budi","indonesia",20}
	filterName :=  filterNames
	filterAge := filterAges
	budi.sayHelloWithFilter("Anjing",10,filterName,filterAge)
}

// var eko Customer
	// eko.Name = "Eko kurniawan"
	// eko.Address = "Indonesia"
	// eko.Age = 30
	
	// fmt.Println(eko.Name)

	// joko := Customer{
	// 	Name: "joko",
	// 	Address: "Indonesia",
	// 	Age: 20,
	// }
	// fmt.Println(joko)