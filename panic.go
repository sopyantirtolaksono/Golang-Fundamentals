package main

import "fmt"

func EndApp() {
	fmt.Println("APLIKASI BERHENTI!")
	message := recover()
	if message != nil {
		fmt.Println("ISI PESAN ERRORNYA ADALAH:", message)
	}
}

func RunApp(err bool) {
	defer EndApp()
	if err {
		panic("APLIKASI ERROR!")
	}
	fmt.Println("APLIKASI BERJALAN!")
}

func main() {
	RunApp(true)
}
