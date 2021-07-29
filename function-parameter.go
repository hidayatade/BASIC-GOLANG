package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {

	firstName := "ade"
	lastName := "hidayat"

	sayHelloTo(firstName, lastName)
	sayHelloTo("bambang", "jamal")
}
