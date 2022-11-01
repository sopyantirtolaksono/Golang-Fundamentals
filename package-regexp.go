package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eki"))
	fmt.Println(regex.MatchString("edo"))

	fmt.Println(regex.FindAllString("eko, eki, edo, edi, ego", -1))

}
