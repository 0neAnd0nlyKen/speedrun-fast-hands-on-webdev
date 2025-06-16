package handler

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins for CORS

	dataJson := map[string]interface{}{
		"data": map[string]interface{}{
			"message": "Hello World!",
		},
	}

	err := json.NewEncoder(w).Encode(dataJson)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func Serve() {
	r := http.NewServeMux()
	r.HandleFunc("GET /hello", HomeHandler)
	r.HandleFunc("POST /hello", HomeHandler)
	http.ListenAndServe(":8080", r)
}
