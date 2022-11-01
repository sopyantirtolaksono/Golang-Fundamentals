package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noKtpSopyan NoKtp = "1234567890"
	var marriedStatus Married = false

	fmt.Println(noKtpSopyan)
	fmt.Println(marriedStatus)
}