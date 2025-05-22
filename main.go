package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Handler for root path "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Gorilla Mux!")
}

// Handler with path variable
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/greet/{name}", GreetHandler).Methods("GET")

	// Start the HTTP server
	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
