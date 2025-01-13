package main

import (
	"fmt"
	"golang_api/database"
	"log"
)

func main() {
	db := database.GetConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, created_at, user_id, device_id FROM user_devices")
	if err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var created_at, user_id, device_id string
		if err := rows.Scan(&id, &created_at, &user_id, &device_id); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		fmt.Printf("ID: %d, CreatedAt: %s, User: %s, Devices: %s\n", id, created_at, user_id, device_id)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error during row iteration: %v", err)
	}
}
