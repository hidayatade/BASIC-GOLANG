package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ade hidayat", "Ade"))
	fmt.Println(strings.Contains("Ade hidayat", "Jamal"))

	fmt.Println(strings.Split("Ade Hidayat", " "))

	fmt.Println(strings.ToLower("Ade Hidayat"))
	fmt.Println(strings.ToUpper("Ade Hidayat"))
	fmt.Println(strings.ToTitle("Ade Hidayat"))

	fmt.Println(strings.Trim("    Ade Hidayat    ", " "))
	fmt.Println(strings.ReplaceAll("Ade Ade Ade", "Ade", "Hidayat"))
}
