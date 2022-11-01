package main

import (
	"belajar-golang-dasar/helper"
	"belajar-golang-dasar/other"
	"fmt"
)

func main() {
	hai := helper.SayHello("Ferry")
	fmt.Println(hai)

	hei := other.SayHello("Sopyan")
	fmt.Println(hei)

	// fmt.Println(helper.Name)
	// fmt.Println(helper.Address)

	goodBye := helper.SayGoodBye("Chabib")
	fmt.Println(goodBye)

	fmt.Println(helper.CreateNewStudent("Sopyan", "Kendal"))
	fmt.Println(helper.UpdateStudent("Ferry", "Semarang"))
}
