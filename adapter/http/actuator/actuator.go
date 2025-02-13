package actuator

import (
	"encoding/json"
	"net/http"
)

type HealthBody struct {
	Status string `json:"status"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	
	var status = HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)
}
