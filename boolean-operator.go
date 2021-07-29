package main

import "fmt"

func main() {

	var nilaiUjian = 90
	var nilaiAbsensi = 80

	// var lulusUjian = nilaiUjian >= 80
	// var lulusAbsensi = nilaiAbsensi >= 90
	// fmt.Println(lulusUjian)
	// fmt.Println(lulusAbsensi)

	// var lulus = lulusUjian && lulusAbsensi

	// fmt.Println(lulus)
	fmt.Println(nilaiUjian >= 80 && nilaiAbsensi >= 80)

}
