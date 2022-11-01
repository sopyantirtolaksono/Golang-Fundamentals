package main

import "fmt"

type BlackList func(string) bool

func RegisterUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}

	RegisterUser("admin", blackList)
	RegisterUser("sopyan", blackList)

	AnonymousFunc := func(name string) string {
		return "Hello, I'am " + name + " & I'am anonymous function!"
	}

	fmt.Println(AnonymousFunc("Fulan"))
}
