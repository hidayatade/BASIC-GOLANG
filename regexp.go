package main

import (
	"fmt"
	"regexp"
)

func main() {

	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eTo"))
	fmt.Println(regex.MatchString("edo"))

	search := regex.FindAllString("eko, eki, eka, eyo, edo", -1)
	fmt.Println(search)
}
