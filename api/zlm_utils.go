package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func HttpResult(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if resp != nil {
		buf, err := json.Marshal(resp)

		if err != nil {
			log.Println("HttpResult json.Marshal error:", err)
			return
		}
		w.Write(buf)
	}
}
