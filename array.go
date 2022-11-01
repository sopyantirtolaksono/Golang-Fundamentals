package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Sopyan"
	names[1] = "Tirto"
	names[2] = "Laksono"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int {
		11,
		10,
		95,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var animals = [4]string {
		"Kucing",
		"Sapi",
		"Ayam",
		"Bebek",
	}
	fmt.Println(animals)

	// var (
	// 	food = "Cake"
	// 	drink = "Milk"
	// )

	var fruits = [4]string {
		"banana",
		"apple",
		"orange",
		"grape",
	}

	fmt.Println(fruits)
}