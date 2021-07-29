package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "mr. " + man.Name
}

func main() {
	ade := Man{"hidayat ade"}
	ade.Married()
	fmt.Println(ade.Name)
}
