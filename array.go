package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "ade"
	names[1] = "hi"
	names[2] = "dayat"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		80,
		90,
		100,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}
