package main

import "fmt"

func main() {
	// var names [3]string

	// names[0] = "eko"
	// names[1] = "kurniawan"
	// names[2] = "khannedy"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	var values = [...]int{
		90,
		80,
		85,
		105,
	}
	values[1] = 200
	fmt.Println(values)
	fmt.Println(len(values))
}