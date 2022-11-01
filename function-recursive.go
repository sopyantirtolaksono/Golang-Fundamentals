package main

import "fmt"

func Recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * Recursive(value-1)
	}
}

func main() {
	recursive := Recursive(10)
	fmt.Println(recursive)
}
