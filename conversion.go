package main

import "fmt"

func main() {
	var int32 int32 = 100000
	var int64 int64 = int64(int32)
	var int8 int8   = int8(int32)

	fmt.Println(int32)
	fmt.Println(int64)
	fmt.Println(int8)

	var name string = "Sopyan"
	var nameConvert string = string(name[0])
	// var nameConvert byte = name[0]
	
	fmt.Println(len(name))
	fmt.Println(nameConvert)
}