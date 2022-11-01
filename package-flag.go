package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Connection host")
	user := flag.String("user", "root", "Connection database user")
	pass := flag.String("pass", "root", "Connection database pass")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*user)
	fmt.Println(*pass)
}
