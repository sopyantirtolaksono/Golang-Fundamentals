package main

import "fmt"

type Filter func(string) string
type FilterVirus func(string) string

func SayHelloWithFilter(name string, filter Filter) {

	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)

}

func SpamFilter(name string) string {

	if name == "Anjing" {
		return "..."
	} else {
		return name
	}

}

func VirusDectection(name string, filterVirus FilterVirus) {

	fmt.Println(filterVirus(name))

}

func SpamFilterVirus(name string) string {

	if name == "babi" {
		return "..."
	} else {
		return "Hello, " + name
	}

}

func main() {

	SayHelloWithFilter("Anjing", SpamFilter)

	VirusDectection("kucing", SpamFilterVirus)

}

// func main() {

// 	slice := []string {"sopyan", "tirto", "laksono"}
// 	fmt.Println(slice[0])

// 	fmt.Println("====================================")

// 	dataMap := map[int] int {
// 		0 : 1,
// 		1 : 2,
// 		2 : 3,
// 		3 : 4,
// 		4 : 5,
// 	}

// 	for _, val := range dataMap {
// 		// fmt.Println(key)
// 		fmt.Println(val)
// 	}

// }
