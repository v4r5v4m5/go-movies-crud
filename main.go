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

// movie struct goes here
type Movie struct {
	ID       string    `json:"id"`       // id of the movie
	Isbn     string    `json:"isbn"`     // rating kind of stuff
	Title    string    `json:"title"`    // title of the movie
	Director *Director `json:"director"` // director struct returned with respect to movies
}

// director struct goes here
type Director struct {
	Firstname string `json:"firstname"` // first name of the director
	Lastname  string `json:"lastname"`  // last name of the director
}

// slice of movies
var movies []Movie

// get list of movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // here we are looking for the movie to be deleted and substituting rest of the elements in that index to be deleted
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// get a particular movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

// create movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// update movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter() // creates a new router

	// lets add some movies
	movies = append(movies, Movie{ID: strconv.Itoa(rand.Intn(100000000)), Isbn: strconv.Itoa(rand.Intn(6969)), Title: "SVSC", Director: &Director{Firstname: "Srikanth", Lastname: "Addala"}})
	movies = append(movies, Movie{ID: strconv.Itoa(rand.Intn(100000000)), Isbn: strconv.Itoa(rand.Intn(6969)), Title: "hey pillagada", Director: &Director{Firstname: "Sameer", Lastname: "Thahir"}})
	movies = append(movies, Movie{ID: strconv.Itoa(rand.Intn(100000000)), Isbn: strconv.Itoa(rand.Intn(6969)), Title: "orange", Director: &Director{Firstname: "Bommarillu", Lastname: "Bhaskar"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")           // get all movies
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // get a specific movie
	r.HandleFunc("/movies", createMovie).Methods("POST")        // create a new movie
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // update an existing movie
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // delete a movie

	// golang server
	fmt.Printf("Strting golang server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
