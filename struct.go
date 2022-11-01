package main

import "fmt"

// Deklarasi struct
type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) SayHello() {
	fmt.Println("Hello, my name is", customer.Name, "and i'am from", customer.Address, "and i'am", customer.Age, "years old")
}

func main() {
	// Pengisian data struct biasa
	var saya Customer
	saya.Name = "Sopyan"
	saya.Address = "Kendal"
	saya.Age = 26

	saya.SayHello()

	// fmt.Println(saya.Name)
	// fmt.Println(saya.Address)
	// fmt.Println(saya.Age)

	// Pengisian data struct literal
	// me := Customer{
	// 	Name:    "Ferry",
	// 	Address: "Semarang",
	// 	Age:     19,
	// }

	// fmt.Println(me)

	// Pengisian data struct literal(Not Recommended)
	// kulo := Customer{"Yohanes", "Yogyakarta", 25}

	// fmt.Println(kulo)
}
