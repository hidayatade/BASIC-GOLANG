package main

import "fmt"

func getGoodBye(names string) string {
	return "bye " + names
}

func main() {

	sayGoodBye := getGoodBye
	result := getGoodBye("hidayat")

	fmt.Println(result)
	fmt.Println(sayGoodBye("ade"))

}
