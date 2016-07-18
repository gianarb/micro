package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"log"
)

type HealtResponse struct {
	Status bool              `json:"status"`
	Info   map[string]string `json:"info"`
}

func health(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := HealtResponse{Status: true}
		httpStatus := 200
		if err := db.Ping(); err != nil {
			res.Status = false
			res.Info = map[string]string{"database": err.Error()}
		}
		c, _ := json.Marshal(res)
		if res.Status == false {
			httpStatus = 500
		}
		log.Println("%s called /health", r.Host)
		w.WriteHeader(httpStatus)
		w.Header().Set("Content-Type", "application/json")
		w.Write(c)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("%s called /", r.Host)
	w.Header().Set("Content-Type", "text/html")
	ifaces, _ := net.Interfaces()
	var ip net.IP
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}
	b, _ := ip.MarshalText()
	res := fmt.Sprintf("<div style='text-align:center;'><p style='font-size:5em'>%s</p><p style='font-size:3em'>v2.0</p></div>", b)
	w.Write([]byte(res))
}

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
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", hello)
	http.HandleFunc("/health", health(db))
	http.ListenAndServe(":8000", nil)
}
