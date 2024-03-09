package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	person := Person{firstName: "John", lastName: "Doe"}

	fmt.Println(person.firstName, person.lastName)
}
