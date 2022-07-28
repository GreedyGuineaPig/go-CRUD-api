package main

// TODO: go fmt main.go before merge to master!

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...) // not efficient but because we don't have DB
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	fmt.Println(id)

	for _, item := range movies {
		fmt.Println(item.Title)
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = len(movies) + 1
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func initializeMovies() {
	movies = nil
	// add some sample movies
	movies = append(movies, Movie{ID: 1, Title: "Titanic", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	movies = append(movies, Movie{ID: 2, Title: "E.T.", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	movies = append(movies, Movie{ID: 3, Title: "Jurassic Park", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	initializeMovies()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/movies", getAllMoviesHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovieHandler).Methods("DELETE")
	r.HandleFunc("/movies/{id}", getMovieHandler).Methods("GET")
	r.HandleFunc("/movies", createMovieHandler).Methods("POST")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
