package main

import (
	//_	"firstGo/database" // tambahhan _ jika inport tidak dipakai / blank idifier
	"firstGo/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
