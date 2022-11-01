package main

import "fmt"
// import "strconv"

func main() {

	fmt.Println("===============================")

	// deklarasi map
	name := map[string] string {
		"sopyan" : "programmer",
		"ferry" : "desainer",
		"Oops" : "salah",
	}

	// menampilkan map - apabila hanya menampilkan map nya tanpa key, maka hasil output akan random/tdak urut
	fmt.Println(name)
	fmt.Println(name["sopyan"])
	fmt.Println(len(name))

	// menghapus map
	delete(name, "Oops")
	fmt.Println(name)
	fmt.Println(len(name))

	fmt.Println("==============================")

	// membuat map baru
	var books map[string]string = make(map[string]string)
	books["title"] = "How to become expert in golang?"
	books["author"] = "sopyan"
	books["ups"] = "salah"

	fmt.Println(books)
	fmt.Println(books["ups"])
	fmt.Println(len(books))

	// hapus data pada map books
	delete(books, "ups")
	fmt.Println(books)
	fmt.Println(books["ups"])
	fmt.Println(len(books))
}