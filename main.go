package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname  string `json:"firstname"`
	Secondname string `json:"secondname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)

	for _, item := range movies {
		if item.ID == parameters["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa((rand.Intn(100000000)))

	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for i, item := range movies {
		if item.ID == parameters["id"] {

			movies = append(movies[:i], movies[i+1:]...)
			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa((rand.Intn(100000000)))

			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)

	for i, item := range movies {
		if item.ID == parameters["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	//the movie slice which we will perform our operations on

	movies = append(movies, Movie{ID: "1", Isbn: "52321", Title: "the 1 movie", Director: &Director{Firstname: "Jhon 1", Secondname: "THE JHON 1"}})
	movies = append(movies, Movie{ID: "2", Isbn: "52323", Title: "the 2 movie", Director: &Director{Firstname: "Jhon 2", Secondname: "THE JHON 2"}})
	movies = append(movies, Movie{ID: "3", Isbn: "52334", Title: "the 3 movie", Director: &Director{Firstname: "Jhon 3", Secondname: "THE JHON 3"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("Starting server at port :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
