package handler

//import local package handler
import (
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /", HomeHandler)
	http.ListenAndServe(":8080", r)
}
