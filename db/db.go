package db

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	_ "github.com/lib/pq"
)

func ConectaBancoDados() *sql.DB {
	//connStr := "user=postgres dbname=digport_loja password=digport host=localhost sslmode=disable"
	dbPass := os.Getenv("DB_PASS")
	connStr := fmt.Sprint("user=postgres dbname=digport_loja password=digport passowrd=", dbPass, " host=localhost sslmode=disable")
	db, err :=sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
