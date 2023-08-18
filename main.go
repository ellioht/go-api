package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

func ContextMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mydata := chi.URLParam(r, "mydata")                          // Get the value from the URL
		ctx := context.WithValue(r.Context(), "mydataValue", mydata) // Add it to the context
		next.ServeHTTP(w, r.WithContext(ctx))                        // Call the next handler
	})
}

func main() {
	// Create a new Router
	r := chi.NewRouter()

	// GET
	r.With(ContextMiddleWare).Get("/api/v1/test/{mydata}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()                               // Get the context
		mydataValue := ctx.Value("mydataValue").(string) // Get the value from the context
		w.Write([]byte(mydataValue))
	})

	// Start the server
	http.ListenAndServe(":8000", r)
}
