package main

import "fmt"

type Addresses struct {
	City, Province, Country string
}

func ChangeCountry(addresses *Addresses) {
	addresses.Country = "Indonesia"
}

func main() {

	address1 := Addresses{"Bekasi", "Jawa Barat", "Indonesia"}
	addrres2 := &address1

	addrres2.City = "Bandung"
	// fmt.Println(address1)
	// fmt.Println(addrres2)

	*addrres2 = Addresses{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(addrres2)

	var address4 = new(Addresses)
	address4.City = "Indonesia"
	fmt.Println(address4)

	var alamat = Addresses{
		City:     "Cirebon",
		Province: "Jawa Barat",
		Country:  "",
	}
	ChangeCountry(&alamat)

	fmt.Println(alamat)
}
