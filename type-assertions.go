package main

import "fmt"

func Random() interface{} {
	return true
}

func main() {
	var result interface{} = Random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println(value, "is string")
	case int:
		fmt.Println(value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
