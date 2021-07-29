package main

import "fmt"

func getHello(name string) string {

	if name == "" {
		return "Hello brow"
	} else {
		return "Hello " + name
	}
}

func main() {

	result := getHello("ade")
	fmt.Println(result)

	fmt.Println(getHello(""))

}
