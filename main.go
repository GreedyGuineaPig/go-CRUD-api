package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "this is a root page")
}

func getAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func initializeMovies() {
	// add some sample movies
	movies = append(movies, Movie{ID: 1, Title: "Titanic", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	movies = append(movies, Movie{ID: 2, Title: "E.T.", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	movies = append(movies, Movie{ID: 3, Title: "Jurassic Park", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/movies", getAllMoviesHandler).Methods("GET")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
