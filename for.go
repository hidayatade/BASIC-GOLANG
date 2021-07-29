package main

import "fmt"

func main() {

	counter := 1

	for counter <= 5 {
		fmt.Println("counter ke: ", counter)
		counter++
	}

	fmt.Println("======== FOR STATEMENT =========")

	for counter2 := 1; counter2 <= 5; counter2++ {
		fmt.Println("counter ke: ", counter2)
	}

	fmt.Println("========== Slice ================")

	slice := []string{"ade", "hi", "dayat"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("========== For Range =============")

	// kalo tidak mau gunakan index, i ganti jadi undescore(_)
	for i, values := range slice {
		fmt.Println("index:", i, "=", values)
	}

	fmt.Println("========== make ===============")

	person := make(map[string]string)

	person["name"] = "ade"
	person["title"] = "software engineer"

	for key, values := range person {
		fmt.Println(key, "=", values)
	}
}
