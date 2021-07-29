package main

import "fmt"

func main() {
	var name string

	name = "ade hidayat"
	fmt.Println(name)

	name = "hidayat ade"
	fmt.Println(name)

	var number int

	number = 25
	fmt.Println(number)

	number = 20
	fmt.Println(number)

	addres := "bekasi"
	fmt.Println(addres)

	addres = "jakarta"
	fmt.Println(addres)

	var (
		firstName = "ade"
		lastName  = "hidayat"
	)
	fmt.Println("nama depan :", firstName)
	fmt.Println("nama belakang :", lastName)
}
