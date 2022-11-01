package main

import "fmt"

func Random() interface{} {
	// return "Hello..."
	// return 100
	return true
}

func main() {
	var result interface{} = Random()
	switch value := result.(type) {
	case string:
		fmt.Println(value, "adalah string")
	case int:
		fmt.Println(value, "adalah integer/number")
	default:
		fmt.Println("Unknown type")
	}
}
