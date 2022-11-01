package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Sopyan Tirto Laksono", "Sopyan"))
	fmt.Println(strings.Contains("Sopyan Tirto Laksono", "Ferry"))

	fmt.Println("=====================================")

	fmt.Println(strings.Split("Sopyan Tirto Laksono", " "))

	fmt.Println("=====================================")

	fmt.Println(strings.ToLower("Sopyan Tirto Laksono"))
	fmt.Println(strings.ToUpper("Sopyan Tirto Laksono"))
	fmt.Println(strings.ToTitle("sopyan tirto laksono"))

	fmt.Println("=====================================")

	fmt.Println(strings.Trim("   Sopyan Tirto Laksono   ", " "))
	fmt.Println(strings.Trim("x   Sopyan Tirto Laksono   x", " "))

	fmt.Println("=====================================")

	fmt.Println(strings.ReplaceAll("Sopyan Chabib Sopyan Sopyan", "Sopyan", "Ferry"))
}
