package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func helloMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloMuxHandler)
	http.ListenAndServe(":8080", r)

}
