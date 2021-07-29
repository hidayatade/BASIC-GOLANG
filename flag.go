package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database localhost")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your database number")

	flag.Parse()

	fmt.Println("HOST: ", *host)
	fmt.Println("USER:", *user)
	fmt.Println("PASS:", *password)
	fmt.Println("NUMBER", *number)
}
