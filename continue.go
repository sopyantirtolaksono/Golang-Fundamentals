package main

import "fmt"

func main() {

	// continue(kondisi dimana jika statement terpenuhi maka akan di skip dan lanjut ke kode berikutnya)
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		
		fmt.Println("Perulangan ke -", i)
	}

}