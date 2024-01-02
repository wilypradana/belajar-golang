package main

import "fmt"

func getfullname()(string, string, int)  {
	return "Eko", "kurniawan", 12
}

func main() {
	firstname, _ := getfullname()
	fmt.Println(firstname)
}