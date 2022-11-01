package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	fmt.Println("=================================")

	utc := time.Date(2022, time.March, 8, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	fmt.Println("=================================")

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "1945-08-17")
	fmt.Println(parse)
}
