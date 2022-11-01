package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b

	fmt.Println(c)

	a += 5
	a += 3
	b -= 1
	c -= 15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var d = 1
	var e = 5
	d++
	e--

	fmt.Println(d)
	fmt.Println(e)

	var f = +100
	var g = -100

	fmt.Println(f)
	fmt.Println(g)

}