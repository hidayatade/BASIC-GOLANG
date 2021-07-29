package main

import "fmt"

type Hasname interface {
	getName() string
}

func sayHellow(hasName Hasname) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

func main() {
	var ade Person
	ade.Name = "ade"

	sayHellow(ade)
}
