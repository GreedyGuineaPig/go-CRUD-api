package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// func TestMain(m *testing.M) {
// 	initializeMovies()
// 	code := m.Run()
// 	os.Exit(code)
// }

func TestRootHandler(t *testing.T) {
	initializeMovies()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	rootHandler(rw, req)
	res := rw.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	rootMessage := "this is a root page"

	assert.NoError(t, err)
	assert.Equal(t, rootMessage, string(data))
}

func TestGetAllMoviesHandler(t *testing.T) {
	initializeMovies()

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/movies", nil)
	getAllMoviesHandler(rw, req)

	res := rw.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	var actualMovies []Movie
	json.Unmarshal(data, &actualMovies)

	var expectedMovies []Movie
	expectedMovies = append(expectedMovies, Movie{ID: 1, Title: "Titanic", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	expectedMovies = append(expectedMovies, Movie{ID: 2, Title: "E.T.", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	expectedMovies = append(expectedMovies, Movie{ID: 3, Title: "Jurassic Park", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})

	assert.NoError(t, err)
	assert.Equal(t, expectedMovies, actualMovies)
}

func TestDeleteMovieHandler1stElement(t *testing.T) {
	initializeMovies()

	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", deleteMovieHandler).Methods("DELETE")

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/movies/1", nil)

	router.ServeHTTP(rw, req)

	res := rw.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	var actualMovies []Movie
	json.Unmarshal(data, &actualMovies)

	var expectedMovies []Movie
	// expectedMovies = append(expectedMovies, Movie{ID: 1, Title: "Titanic", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	expectedMovies = append(expectedMovies, Movie{ID: 2, Title: "E.T.", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	expectedMovies = append(expectedMovies, Movie{ID: 3, Title: "Jurassic Park", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expectedMovies, actualMovies)
}

func TestDeleteMovieHandler2ndElement(t *testing.T) {
	initializeMovies()

	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", deleteMovieHandler).Methods("DELETE")

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/movies/2", nil)

	router.ServeHTTP(rw, req)

	res := rw.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	var actualMovies []Movie
	json.Unmarshal(data, &actualMovies)

	var expectedMovies []Movie
	expectedMovies = append(expectedMovies, Movie{ID: 1, Title: "Titanic", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	// expectedMovies = append(expectedMovies, Movie{ID: 2, Title: "E.T.", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	expectedMovies = append(expectedMovies, Movie{ID: 3, Title: "Jurassic Park", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})

	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expectedMovies, actualMovies)
}

func TestGetMovieHandler(t *testing.T) {
	initializeMovies()

	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", getMovieHandler).Methods("GET")

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/movies/1", nil)

	router.ServeHTTP(rw, req)

	res := rw.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	var actualMovie Movie
	json.Unmarshal(data, &actualMovie)

	expectedMovie := Movie{ID: 1, Title: "Titanic", Director: &Director{Firstname: "James", Lastname: "Cameron"}}

	assert.NoError(t, err)
	assert.Equal(t, expectedMovie, actualMovie)
}

func TestCreateMovieHandler(t *testing.T) {
	initializeMovies()
	movie := Movie{ID: 1, Title: "Jaws", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}}
	byteMovie, _ := json.Marshal(movie)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/movies", bytes.NewReader(byteMovie))

	createMovieHandler(rw, req)

	res := rw.Result()

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	var actualMovies []Movie
	json.Unmarshal(data, &actualMovies)

	movie.ID = 4
	var expectedMovies []Movie
	expectedMovies = append(expectedMovies, Movie{ID: 1, Title: "Titanic", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	expectedMovies = append(expectedMovies, Movie{ID: 2, Title: "E.T.", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	expectedMovies = append(expectedMovies, Movie{ID: 3, Title: "Jurassic Park", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	expectedMovies = append(expectedMovies, movie)

	assert.NoError(t, err)
	assert.Equal(t, expectedMovies, actualMovies)
}

func TestUpdateMovieHandler(t *testing.T) {
	initializeMovies()

	router := mux.NewRouter()
	router.HandleFunc("/movies/{id}", updateMovieHandler).Methods("PUT")

	// + remastered to the title
	movie := Movie{ID: 1, Title: "Titanic remastered", Director: &Director{Firstname: "James", Lastname: "Cameron"}}
	byteMovie, _ := json.Marshal(movie)

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/movies/1", bytes.NewReader(byteMovie))

	router.ServeHTTP(rw, req)

	res := rw.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	var actualMovies []Movie
	json.Unmarshal(data, &actualMovies)

	var expectedMovies []Movie
	expectedMovies = append(expectedMovies, Movie{ID: 1, Title: "Titanic remastered", Director: &Director{Firstname: "James", Lastname: "Cameron"}})
	expectedMovies = append(expectedMovies, Movie{ID: 2, Title: "E.T.", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})
	expectedMovies = append(expectedMovies, Movie{ID: 3, Title: "Jurassic Park", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})

	assert.NoError(t, err)
	assert.Equal(t, expectedMovies, actualMovies)
}
