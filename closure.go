package main

import "fmt"

func main() {
	name := "Jamal"
	counter := 0
	fmt.Println(counter)

	increment := func() {
		// name := "Bambang"
		name = "Bambang"
		fmt.Println("increment")
		counter++
		// fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)

}
