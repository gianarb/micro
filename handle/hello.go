package handle

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func Hello(version string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("%s called /", r.Host)
		w.Header().Add("Content-Type", "text/html")
		t := template.Must(template.New("foo").Parse(`<h2>Hello, I am micro {{.Version}}</h2><ul><li><a href="/net">Network</a><li><a href="/fs">File Systems</a></li><li><a href="/dns">resolv.conf</a></li><li><a href="/env">Env Vars</a></li><li><a href="/health">health</a></li></ul><p>Date {{.Date}}</p>`))
		t.Execute(w, map[string]string{
			"Date":    time.Now().String(),
			"Version": version,
		})
	}
}
