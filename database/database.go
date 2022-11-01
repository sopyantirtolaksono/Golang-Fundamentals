package database

import "fmt"

var connection string

func init() {
	fmt.Println("Calling database...")
	connection = "MYSQL"
}

func GetDatabase() string {
	return connection
}
