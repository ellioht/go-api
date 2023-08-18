package main

import (
	"context"
	"net/http"

	"fmt"

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

	// GET with path parameter
	r.With(ContextMiddleWare).Get("/api/v1/test/{mydata}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()                               // Get the context
		mydataValue := ctx.Value("mydataValue").(string) // Get the value from the context
		w.Write([]byte(mydataValue))                     // Write it to the response
	})

	// GET with query parameter
	r.Get("/api/v1/puzzle", func(w http.ResponseWriter, r *http.Request) {
		numeral := r.URL.Query().Get("numeral") // Get the value of the "mydata" query parameter

		// Convert the Roman numeral string to an integer
		romanValues := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000, 'i': 1, 'v': 5, 'x': 10, 'l': 50, 'c': 100, 'd': 500, 'm': 1000}
		total := 0
		prevValue := 0
		for _, c := range numeral {
			value := romanValues[c]
			if value > prevValue {
				total += value - 2*prevValue
			} else {
				total += value
			}
			prevValue = value
		}

		fmt.Println(total)

		fmt.Fprintf(w, "The Roman numeral value of %s as an integer is %d", numeral, total)
	})

	// Start the server
	http.ListenAndServe(":8000", r)
}
