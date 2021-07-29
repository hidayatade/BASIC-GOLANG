package main

import "fmt"

func getFullname() (string, string, string) {
	return "Ade", "Hi", "dayat"
}

func main() {

	// jika tidak mau guankan vaiable maka ganti denga underscroe(_)
	// firstName, _, lastName := getFullname()

	firstName, middleName, lastName := getFullname()
	fmt.Println(firstName, middleName, lastName)
	fmt.Println(lastName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
