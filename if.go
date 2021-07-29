package main

import "fmt"

func main() {

	name := "HidaYAT"

	if name == "Hidayat" {
		fmt.Println("Hello ade")
	} else if name == "Joko" {
		fmt.Println("Hello joko")
	} else if name == "Bambang" {
		fmt.Println("Hello bambang")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
