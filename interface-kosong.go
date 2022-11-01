package main

import "fmt"

func Ups(i int, apapun interface{}) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		if apapun == "salah" {
			return false
		}
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(3, "salah")
	fmt.Println(data)
}
