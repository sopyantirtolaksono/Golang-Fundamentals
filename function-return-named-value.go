package main

import "fmt"

func GetNames() (firstname string, middlename string, lastname string) {

	firstname = "Sopyan"
	middlename = "Tirto"
	lastname = "Laksono"

	return

}

func JoinAddress() (addr1, addr2, addr3 string) {
	addr1 = "Tlahab"
	addr2 = "Kendal"
	addr3 = "Jawa Tengah Indonesia"

	return
}

func main() {

	firstname, middlename, lastname := GetNames()

	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)

	a, b, c := JoinAddress()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
