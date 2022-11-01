package main

import "fmt"

// deklarasi function dengan return valuenya(satu return value)
func sayHello(name string) string {
	if name == "" {
		return "Hello, Bro!"
	} else {
		return "Hello, " + name
	}
}

func GuestStar(name string) string {
	if name == "" {
		return "Nothing Guest Star!"
	} else {
		return "Hello " + name
	}
}

func main() {
	var result string = "ferry"

	fmt.Println(sayHello(result))
	fmt.Println(sayHello(""))

	fmt.Println(GuestStar(""))
	fmt.Println(GuestStar("Chabib"))
}
