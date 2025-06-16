package handler

import (
	"fmt"
	"net/http"
)

var username string = ""

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // for CORS if needed

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "ParseForm() error", http.StatusBadRequest)
			return
		}
		username = r.FormValue("username")
	}
	fmt.Fprintf(w, "<h1>Hello %s</h1>", username)
}
func ServeUsers() {
	r := http.NewServeMux()
	r.HandleFunc("/", GetUsers)
	http.ListenAndServe(":8080", r)
}
