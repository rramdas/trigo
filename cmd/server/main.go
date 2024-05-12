package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rramdas/trigo/api"
)

var graph *api.Graph

var once sync.Once

func GetGraph() *api.Graph {
	once.Do(func() {
		graph = &api.Graph{}
	})
	return graph
}

func triplePutHandler(w http.ResponseWriter, r *http.Request) {
	// if the graph is not found, create a new one
	// if the graph is found, use the existing one
	// if the graph is found, but the triple is not found, add the triple
	// if the graph is found, and the triple is found, update the triple
	// if the graph is found, and the triple is found, and the triple is the same, do nothing
	// if the graph is found, and the triple is found, and the triple is different, update the triple
}

func helloMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloMuxHandler)
	http.ListenAndServe(":8080", r)

}
