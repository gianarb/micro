package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"log"
)

type netResponse struct {
	Hostname string
	IP       net.IP
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("%s called /health", r.Host)
	w.Header().Set("Content-Type", "application/json")
	c, _ := json.Marshal(map[string]bool{"status": true})
	w.Write(c)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	t := template.Must(template.New("foo").Parse(`<h2>Hello, I am micro</h2><ul><li><a href="/net">Network</a><li><a href="/fs">File Systems</a></li><li><a href="/dns">resolv.conf</a></li><li><a href="/env">Env Vars</a></li><li><a href="/health">health</a></li></ul><p>Date {{.Date}}</p>`))
	t.Execute(w, map[string]string{
		"Date": time.Now().String(),
	})
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
	http.HandleFunc("/", hello)
	http.HandleFunc("/env", env)
	http.HandleFunc("/dns", dns)
	http.HandleFunc("/health", health)
	http.HandleFunc("/net", network)
	http.HandleFunc("/fs", fs)
	http.ListenAndServe(":8000", nil)
}
