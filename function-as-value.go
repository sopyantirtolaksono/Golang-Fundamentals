package main

import "fmt"

func GetGoodBye(name string) string {

	return "Good Bye " + name + "! See you next time."

}

func GetHello(name string) string {

	return "Hello " + name + ", nice to meet you!"

}

func main() {

	sayGoodBye := GetGoodBye
	result := sayGoodBye("Fulan")
	fmt.Println(result)

	fmt.Println(GetHello("Ferry"))
	outputGetHello := GetHello("Chabib")
	fmt.Println(outputGetHello)

}
