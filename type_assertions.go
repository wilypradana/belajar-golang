package main

import "fmt"

func Random() any {
	return nil
}

func main() {
	 result  := Random()
	 
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default: 
	fmt.Println("UNKNOWN")
	}
}