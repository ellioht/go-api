package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// Create a new Router
	router := chi.NewRouter()

	// GET
	router.Get("/api/v1/test/{mydata}", func(w http.ResponseWriter, r *http.Request) {
		message := chi.URLParam(r, "mydata")
		w.Write([]byte(message))
	})

	// Start the server
	http.ListenAndServe(":8000", router)
}
