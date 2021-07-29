package main

import "fmt"

type Filter func(string) string

func sayHelloFilter(name string, filter Filter) {
	filterNames := filter(name)
	fmt.Println("Hello ", filterNames)
}

func spamFilter(name string) string {
	if name == "Anjay" {
		return "......"
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("Ade", spamFilter)
	sayHelloFilter("Anjay", spamFilter)
}
