package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments: ", args)

	// diketik manual di terminal / go run os.go ade hidayat
	fmt.Println(args[1])
	fmt.Println(args[2])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("HOSTNAME: ", name)
	} else {
		fmt.Println("ERROR", err.Error())
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
