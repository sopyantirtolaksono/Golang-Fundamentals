package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{City: "Semarang", Province: "Jateng", Country: "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Kendal"

	*address2 = Address{City: "Jakarta", Province: "DKI Jakarta", Country: "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Bandung"
	// address4 = &Address{City: "Bandung"}
	// var address4 *Address = &Address{City: "Bandung", Province: "Jabar", Country: "Indonesia"}
	fmt.Println(address4)

	alamat := Address{
		City:     "Solo",
		Province: "Jateng",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)

	fmt.Println(alamat)
}
