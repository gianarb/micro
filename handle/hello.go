package handle

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
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
