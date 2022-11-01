package main

import "fmt"

func EndApp() {
	fmt.Println("APLIKASI BERHENTI!")
}

func main() {
	defer EndApp()
	fmt.Println("APLIKASI BERJALAN!")
}
