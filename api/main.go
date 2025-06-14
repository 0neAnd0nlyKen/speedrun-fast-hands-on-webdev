package handler

//import local package handler
import (
	"net/http"
)

func Serve() {
	r := http.NewServeMux()
	r.HandleFunc("GET /", HomeHandler)
	http.ListenAndServe(":8080", r)
}
