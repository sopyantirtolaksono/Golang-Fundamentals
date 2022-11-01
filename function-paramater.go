package main

import "fmt"

// deklarasi function dengan parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}

func ShowFullName(firstName string, lastName string) {
	fmt.Println("My first name is", firstName, "& my last name is", lastName)
}

// function utama
func main() {
	var firstName string = "sopyan"

	sayHelloTo(firstName, "tirto laksono")
	sayHelloTo("ferry", "irawan")

	ShowFullName("Ferry", "Irawan")
}
