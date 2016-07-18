package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gianarb/micro/handle"

	_ "github.com/go-sql-driver/mysql"

	"log"
)

func main() {
	log.Println("Micro started")
	username := os.Getenv("MYSQL_USERNAME")
	if username == "" {
		log.Fatal("MYSQL_USERNAME is required")
	}

	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		log.Fatal("MYSQL_PASSWORD is required")
	}

	mySqlAddr := os.Getenv("MYSQL_ADDR")
	if mySqlAddr == "" {
		log.Fatal("MYSQL_ADDR is required")
	}

	log.Println("Initializing database connection")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/micro", username, password, mySqlAddr)
	_, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handle.Hello)
	http.HandleFunc("/health", handle.Health(username, password, mySqlAddr))
	http.ListenAndServe(":8000", nil)
}
