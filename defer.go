package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runAplication(value int) {

	defer logging()
	fmt.Println("Run aplication")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runAplication(10)
}
