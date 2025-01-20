package main

import (
	tools "dms/tools"
	"fmt"
	"log"
)

func main() {
	// Example usage
	password := "supersecretpassword"

	// Hash the password
	hashedPassword, err := tools.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hashed password:", hashedPassword)

	// Verify the password
	// password = "thisisthewrong password"
	match, err := tools.VerifyPassword(password, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	if match {
		fmt.Println("The password matches!")
	} else {
		fmt.Println("The password does not match.")
	}
}
