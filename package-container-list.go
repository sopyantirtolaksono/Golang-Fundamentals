package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Sopyan")
	data.PushBack("Tirto")
	data.PushBack("Laksono")
	data.PushFront("Ferry")
	data.PushFront("Chabib")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	fmt.Println("=======================")

	fmt.Println(data.Front().Next().Next().Value)
	fmt.Println(data.Back().Prev().Value)

	fmt.Println("=======================")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("=======================")

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
