package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()
	data.PushBack("ade")
	data.PushBack("hi")
	data.PushBack("dayat")
	data.PushFront("Bambang")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("=================================")

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

}
