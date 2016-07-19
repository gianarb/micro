package handle

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type healtResponse struct {
	Status bool              `json:"status"`
	Info   map[string]string `json:"info"`
}

func Health(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := healtResponse{Status: true}
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
