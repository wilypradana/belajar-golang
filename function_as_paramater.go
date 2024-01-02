package main

import "fmt"

									
func sayHelloWithFilter(name string, filter func(string)string)  {
	fmt.Println("Hello", filter(name))
}

// ini akan terafiliasi dengan filter func di paramater sayhelloWithFilter
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}else{
		return name
	}
}

func main() {
	udahDifilter := spamFilter
	sayHelloWithFilter("Anjing", udahDifilter)
}