package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
)

func ContextMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mydata := chi.URLParam(r, "mydata")                          // Get the value from the URL
		ctx := context.WithValue(r.Context(), "mydataValue", mydata) // Add it to the context
		next.ServeHTTP(w, r.WithContext(ctx))                        // Call the next handler
	})
}

func main() {

	// Open or create the BoltDB database file
	db, err := bolt.Open("mydb.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
		// Loop through the characters in the string and add the values together
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

	r.Post("/api/v1/pokemon", func(w http.ResponseWriter, r *http.Request) {

		// Random number
		source := rand.NewSource(time.Now().UnixNano())
		generator := rand.New(source)
		offset := generator.Intn(1118)

		// Make a GET request to the PokeAPI and return a random Pokemon
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?limit=1&offset=%d", offset)
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Parse the JSON response and extract the Pokemon name
		var data struct {
			Results []struct {
				Name string `json:"name"`
			} `json:"results"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pokemonName := data.Results[0].Name

		// Open the BoltDB database file
		db, err := bolt.Open("my.db", 0600, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Create a new bucket for the Pokemon names
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte("pokemon"))
			return err
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert the new Pokemon name into the bucket
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("pokemon"))
			err := b.Put([]byte(pokemonName), []byte("1"))
			return err
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the Pokemon name to the response
		fmt.Fprintf(w, "Random Pokemon: %s", pokemonName)
	})

	// GET all pokemon from the database
	r.Get("/api/v1/pokemon", func(w http.ResponseWriter, r *http.Request) {

		// Open the BoltDB database file
		db, err := bolt.Open("my.db", 0600, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Retrieve all the Pokemon names from the bucket
		var pokemonNames []string
		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("pokemon"))
			c := b.Cursor()
			for k, _ := c.First(); k != nil; k, _ = c.Next() {
				pokemonNames = append(pokemonNames, string(k))
			}
			return nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the Pokemon names to the response
		json.NewEncoder(w).Encode(pokemonNames)
	})

	// POST a word from body to sql db
	r.Post("/api/v1/word", func(w http.ResponseWriter, r *http.Request) {

		// Open or create the SQLite database
		dbsql, err := sql.Open("sqlite3", "mydatabase.db")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dbsql.Close()

		// Create the words table
		_, err = dbsql.Exec(`
			CREATE TABLE IF NOT EXISTS words (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				word TEXT
			)
			`)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Parse the JSON request body
		var data struct {
			Word string `json:"word"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Insert the new word into the database
		_, err = dbsql.Exec("INSERT INTO words (word) VALUES (?)", data.Word)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the word to the response
		fmt.Fprintf(w, "Word: %s", data.Word)

	})

	// GET all words from the SQL database
	r.Get("/api/v1/word", func(w http.ResponseWriter, r *http.Request) {

		// Open or create the SQLite database
		dbsql, err := sql.Open("sqlite3", "mydatabase.db")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dbsql.Close()

		// Retrieve all the words from the database
		var words []string
		rows, err := dbsql.Query("SELECT word FROM words")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var word string
			if err := rows.Scan(&word); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			words = append(words, word)
		}

		// Write the words to the response
		json.NewEncoder(w).Encode(words)
	})

	// Start the server
	http.ListenAndServe(":8000", r)
}
