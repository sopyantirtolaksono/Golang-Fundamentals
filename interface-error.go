package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	output, err := Pembagi(10, 0)
	if err == nil {
		fmt.Println(output)
	} else {
		fmt.Println(err)
	}
}
