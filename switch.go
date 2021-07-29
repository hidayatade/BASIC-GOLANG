package main

import "fmt"

func main() {

	name := "jamallludinn"

	switch name {
	case "Ade":
		fmt.Println("Hello ade")
	case "Bambang":
		fmt.Println("Hello Bambang")
	case "Jamal":
		fmt.Println("Hello Jamal")
	default:
		fmt.Println("kenalan dong")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	length := len(name)

	switch {
	case length > 8:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}

}
