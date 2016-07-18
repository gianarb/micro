package handle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type healtResponse struct {
	Status bool              `json:"status"`
	Info   map[string]string `json:"info"`
}

func Health(username string, password string, addr string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := healtResponse{Status: true}
		httpStatus := 200
		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/micro", username, password, addr)
		ddb, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}
		if err := ddb.Ping(); err != nil {
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
