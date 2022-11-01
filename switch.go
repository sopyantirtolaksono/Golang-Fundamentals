package main 

import "fmt"

func main() {

	var name string = "sopyan"

	// switch dengan statement
	switch name {
	case "sopyan":
		fmt.Println("Halo,", name)
	case "ferry":
		fmt.Println("Halo,", name)
	default:
		fmt.Println("Halo boleh kenalan?")
	}

	fmt.Println("===================================")

	// switch dengan statement dan short declaration
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang!")
	case false:
		fmt.Println("Nama terlalu pendek!")
	}

	fmt.Println("====================================")

	// switch tanpa statement/posisi statement berada pada case
	var length int = len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang!")
	case length > 5:
		fmt.Println("Nama lumayan panjang!")
	default:
		fmt.Println("Nama sudah benar!")
	}

}