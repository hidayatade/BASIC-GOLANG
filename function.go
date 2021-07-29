package main

import "fmt"

func sayHello() {
	fmt.Println("hello go lang")
}

func main() {
	sayHello()

	for i := 0; i <= 10; i++ {
		fmt.Print(i, " ")
		sayHello()
	}
}
