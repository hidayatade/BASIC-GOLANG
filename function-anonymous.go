package main

import "fmt"

type Blacklist func(name string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	blacklist := func(name string) bool {
		return name == "Admin"
	}

	registerUser("Admin", blacklist)
	registerUser("Ade", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("Hidayat", func(name string) bool {
		return name == "root"
	})
}
