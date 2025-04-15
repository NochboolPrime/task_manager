package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

func ConnectPostgres() {
	connStr := os.Getenv("POSTGRES_CONN")
	var err error
	PostgresDB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL: %v", err)
	}
	if err = PostgresDB.Ping(); err != nil {
		log.Fatalf("Ошибка проверки соединения PostgreSQL: %v", err)
	}
	log.Println("Успешное подключение к PostgreSQL")
}
