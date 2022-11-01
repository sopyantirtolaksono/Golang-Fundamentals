package main

import "fmt"

func main() {

	// break(kondisi dimana jika statement terpenuhi maka kode akan berhenti dijalankan semua, termasuk kode yang ada dibawahnya)
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}

		fmt.Println("Perulangan ke-", i)
	}

}