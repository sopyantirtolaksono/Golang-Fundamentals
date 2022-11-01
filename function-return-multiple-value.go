package main

import "fmt"

// deklarasi function dengan multiple return value(value kembaliannya lebih dari satu)
func getFullName() (string, string, string) {
	return "sopyan", "tirto", "laksono"
}

func MixWord(word1, word2, word3 string) (string, string, string) {
	return "Mixed word is " + word1, word2, word3
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	fmt.Println(MixWord("SAYA", "MAKAN", "MARTABAK"))
	_, _, word3 := MixWord("SAYA", "MINUM", "KOPI")
	fmt.Println(word3)
}
