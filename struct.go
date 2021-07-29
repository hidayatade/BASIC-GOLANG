package main

import "fmt"

type Costumer struct {
	Name, Address string
	Age           int
}

func (costumer Costumer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", costumer.Name)
}

func (a Costumer) sayHuuu() {
	fmt.Println("Huuuu", a.Name)
}

func main() {
	var Ade Costumer
	Ade.Name = "Ade"
	Ade.Address = "Bekasi"
	Ade.Age = 23

	Ade.sayHi("Bambang")
	Ade.sayHuuu()

	// fmt.Println(Ade)
	// fmt.Println(Ade.Name)
	// fmt.Println(Ade.Address)
	// fmt.Println(Ade.Age)

	// Jamal := Costumer{
	// 	Name:    "Jamal",
	// 	Address: "Cirebon",
	// 	Age:     25,
	// }
	// fmt.Println(Jamal)

	// Bambang := Costumer{"Bambang", "Bandung", 27}
	// fmt.Println(Bambang)
}
