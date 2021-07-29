package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {

	firstName = "Ade"
	middleName = "Hi"
	lastName = "Dayat"

	return
}

func main() {
	first, middle, last := getCompleteName()

	fmt.Println(first, middle, last)
	fmt.Println(first)
	fmt.Println(middle)
	fmt.Println(last)
}
