package main

import "fmt"

func main() {

	// perulangan counter standart
	var counter int = 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println("===============================")

	// perulangan pada data slice dengan linear time
	slice := []string {"sopyan", "programmer"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	} 

	fmt.Println("===============================")

	// perulangan pada data slice dengan for range(mengembalikan 2 data[index, value])
	for index, value := range slice {
		fmt.Println("Index =", index, "dan Value =", value)
	}

	fmt.Println("===============================")

	// perulangan pada data map baru dengan for range(mengembalikan 2 data[key, value])
	newMap := make(map[string] string)
	newMap["nama"] = "sopyan"
	newMap["hobby"] = "traveling"

	for _, value := range newMap {
		fmt.Println("Keynya = ?", "dan Valuenya =", value)
	}	

}