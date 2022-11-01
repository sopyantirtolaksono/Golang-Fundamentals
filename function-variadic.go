package main

import "fmt"

func SumAll(number ...int) int {

	total := 0

	for _, value := range number {
		total = total + value
	}

	return total

}

func SumAllNumber(number ...uint64) uint64 {
	var total uint64

	for _, value := range number {
		total += value
	}

	return total
}
func main() {

	total := SumAll(2, 2, 2, 2, 2)
	fmt.Println(total)

	slice := []int{5, 5, 5, 5, 5}
	total = SumAll(slice...)
	fmt.Println(total)

	fmt.Println(SumAllNumber(2, 2, 2, 2))
	slice2 := []uint64{10, 10, 10, 10, 10}
	fmt.Println(SumAllNumber(slice2...))
}
