package main

import (
	"fmt"
)

func main() {
	var name string = "chabib"

	if name == "sopyan" {
		fmt.Println("Helo,", name)
	} else if name == "ferry" {
		fmt.Println("Helo,", name)
	} else if name == "yohanes" {
		fmt.Println("Helo,", name)
	} else {
		fmt.Println("Helo, boleh kenalan?")
	}

	fmt.Println("==========================")

	// lengthName := len(name)
	if lengthName := len(name); lengthName > 5 {
		fmt.Println("Nama terlalu panjang!")
	} else {
		fmt.Println("Nama sudah benar!")
	}
}