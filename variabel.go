package main

import "fmt"

func main() {
	var name string = "Sopyan"
	fmt.Println(name)
	name = "Sopyan Tirto Laksono"
	fmt.Println(name)

	var nim = 21140879
	fmt.Println(nim)
	nim = 1234567890
	fmt.Println(nim)

	male := true
	fmt.Println(male)
	male = false
	fmt.Println(male)

	var (
		firstName 	= "Sopyan"
		lastName  	= "Tirto Laksono"
		age       	= 26
		programmer 	= true
	)

	fmt.Println("This is my first name =", firstName)
	fmt.Println("This is my last name =", lastName)
	fmt.Println("This is my age =", age)
	fmt.Println("Programmer is my hobby =", programmer)
}