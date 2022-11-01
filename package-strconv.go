package main

import (
	"fmt"
	"strconv"
)

func main() {
	parseBool, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(parseBool)
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println("===============================")

	parseFloat, err := strconv.ParseFloat("3.14", 64)
	if err == nil {
		fmt.Println(parseFloat)
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println("===============================")

	parseInt, _ := strconv.ParseInt("1000000", 10, 64)
	fmt.Println(parseInt)

	fmt.Println("===============================")

	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool)

	// fmt.Println("===============================")

	// formatFloat := strconv.FormatFloat(3.14, 10, 2, 64)
	// fmt.Println(formatFloat)

	fmt.Println("===============================")

	formatInt := strconv.FormatInt(777, 10)
	fmt.Println(formatInt)

	fmt.Println("===============================")

	itoa := strconv.Itoa(1000)
	fmt.Println(itoa)

	fmt.Println("===============================")

	atoi, _ := strconv.Atoi("1000")
	fmt.Println(atoi)

}
