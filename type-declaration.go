package main

import "fmt"

func main() {

	type noKtp string
	type merried bool

	var noKtpAde noKtp = "1234567789"
	var statusMerried merried = false

	fmt.Println(noKtpAde)
	fmt.Println(statusMerried)

}
