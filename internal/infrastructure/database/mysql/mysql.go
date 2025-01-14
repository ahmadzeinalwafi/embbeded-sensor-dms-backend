package MySQLConnector

import (
	"database/sql"
	"golang_api/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	cfg := config.LoadConfig()

	dsn := cfg.GetString("DB_USER") + ":" + cfg.GetString("DB_PASSWORD") +
		"@tcp(" + cfg.GetString("DB_HOST") + ":" + cfg.GetString("DB_PORT") +
		")/" + cfg.GetString("DB_NAME") + "?parseTime=true"

	log.Printf("DSN: %s", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
