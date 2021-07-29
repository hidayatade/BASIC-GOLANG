package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Ade",
		"address": "Bekasi",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)

	book["name"] = "belajar go"
	book["author"] = "ade"
	book["wrong"] = "ups"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))

}
