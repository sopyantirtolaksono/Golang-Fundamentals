package main 

import "fmt"

// deklarasi function tanpa parameter
func sayHello() {
	fmt.Println("Hello World!")
}

// function utama
func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}