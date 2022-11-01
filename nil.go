package main

import "fmt"

/**
Nil sendiri hanya bisa digunakkan dibeberapa tipe data, yaitu: slice, map, function, interface, pointer, channel
*/
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Sopyan")
	fmt.Println(person)

	var guest map[string]string = NewMap("")
	if guest == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(guest)
	}
}
