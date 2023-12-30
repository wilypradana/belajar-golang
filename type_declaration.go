package main

import "fmt"

func main() {
	type noKtp string

	var ktpEko noKtp = "111111111"
	
	var contoh string = "2222222"
	var contohKtp = noKtp(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohKtp)
}