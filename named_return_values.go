package main

import "fmt"
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "eko"
	middleName = "kurniawan"
	lastName = "khannedy"
	return firstName, middleName, lastName
}

func main() {
	a,b,c := getCompleteName()
	fmt.Println(a,b,c)

}