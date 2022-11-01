package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  uint
}

type Number struct {
	Number int
}

type UserSlice []User

type NumberSlice []Number

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func (number NumberSlice) Len() int {
	return len(number)
}

func (number NumberSlice) Less(i, j int) bool {
	return number[i].Number < number[j].Number
}

func (number NumberSlice) Swap(i, j int) {
	number[i], number[j] = number[j], number[i]
}

func main() {
	users := []User{
		{"Sopyan", 26},
		{"Ferry", 19},
		{"Chabib", 25},
		{"Dea", 15},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)

	fmt.Println("=======================================")

	numb := []Number{
		{10},
		{5},
		{4},
		{7},
		{9},
	}

	fmt.Println(numb)
	sort.Sort(NumberSlice(numb))
	fmt.Println(numb)

}
