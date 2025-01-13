package main

import (
	"fmt"
	"golang_api/config"
)

func main() {
	config := config.LoadConfig()

	someValue := config.GetString("DATABASE_HOST") // Replace with an actual key
	fmt.Println("Value of some_key:", someValue)
}
