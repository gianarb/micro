package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gianarb/micro/handle"
	_ "github.com/go-sql-driver/mysql"
)

const VERSION = "2.1.0"

type netResponse struct {
	Hostname string
	IP       net.IP
}

func env(w http.ResponseWriter, r *http.Request) {
	for _, e := range os.Environ() {
		w.Write([]byte(fmt.Sprintf("<p>%s</p>", e)))
	}
}

func network(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	va := netResponse{}
	hostname, _ := os.Hostname()
	va.Hostname = hostname
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
	va.IP = ip
	t := template.Must(template.New("foo").Parse(`<p><b>Hostname: </b>{{.Hostname}}</p><p><b>IP: </b>{{.IP}}</p>`))
	t.Execute(w, va)
}

func fs(w http.ResponseWriter, r *http.Request) {
	fs, err := ioutil.ReadFile("/etc/mtab")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(fs)
}

func dns(w http.ResponseWriter, r *http.Request) {
	resolv, err := ioutil.ReadFile("/etc/resolv.conf")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(resolv)
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
	_, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handle.Hello(VERSION))
	http.HandleFunc("/health", handle.Health(username, password, mySqlAddr))
	http.HandleFunc("/env", env)
	http.HandleFunc("/dns", dns)
	http.HandleFunc("/net", network)
	http.HandleFunc("/fs", fs)
	http.ListenAndServe(":8000", nil)
}
