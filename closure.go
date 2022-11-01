package main

import "fmt"

func main() {
	name := "Sopyan"
	fmt.Println(name)

	closure := func() {
		name := "Ferry"
		fmt.Println(name)
	}

	closure()
}
