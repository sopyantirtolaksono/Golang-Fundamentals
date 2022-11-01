package main

import "fmt"

func main() {
	var months = [...]string {
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	// var slice1 = months[4:7]

	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// slice1[0] = "bukan mei"
	// fmt.Println(slice1)

	// var slice2 = months[11:]
	// var slice3 = append(slice2, "sopyan")

	// fmt.Println(slice2)
	// fmt.Println(slice3)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "sopyan"
	newSlice[1] = "tirto-laksono"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)
	
}