package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is a root directory")
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	// add some sample movies
	movies = append(movies, Movie{ID: "1", Title: "Titanic", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	movies = append(movies, Movie{ID: "2", Title: "E.T.", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	movies = append(movies, Movie{ID: "3", Title: "Jurassic Park", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})

	r.HandleFunc("/", rootHandler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
