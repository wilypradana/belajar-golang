package main

import "fmt"


type HasName interface {
 GetName() string	
}
func SayHello(Value HasName)  {
	fmt.Println("Hello", Value.GetName())
}

type Person struct{
	Name string
}
type Animal struct{
	Name string
}
func (person Person) GetName() string {
	return person.Name
}
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Eko"}
	SayHello(person)
	
	animal := Animal{Name: "kucing"}
	SayHello(animal)
	
}