[![Portfolio](https://img.shields.io/badge/CTF/Writeups-Jayavamsi-orange?style=for-the-badge&logo=netlify)](https://jayavamsi.netlify.app/)
[![Drop](https://img.shields.io/badge/Drop-Email-red?style=for-the-badge&logo=proton)](mailto:shamsi9@protonmail.com)

# go movies crud
This is a movies API project built using go-lang. It offers CRUD operations where a user can Create, Get/Read, Update and Delete a movie. This doesn't use any database for backend, hence the seeding is done initially and movie data stored in slices via structs and pointer references.

## structs
* `movie` struct goes here
```go
type Movie struct {
	ID       string    `json:"id"`       // id of the movie
	Isbn     string    `json:"isbn"`     // rating kind of stuff
	Title    string    `json:"title"`    // title of the movie
	Director *Director `json:"director"` // director struct returned with respect to movies
}
```
* `director` struct goes here
```go
type Director struct {
	Firstname string `json:"firstname"` // first name of the director
	Lastname  string `json:"lastname"`  // last name of the director
}
```

## slice
slice of movies used for storing multiple movies
```go
var movies []Movie
```

## functions
* `getMovies` function is used for returning list of movies using json marshalling.
```go
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
```
* `deleteMovie` function is used for deleting a movie and returning using json marshalling.
```go
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
```
* `getMovie` function is used for getting a particular movie and returning using json marshalling.
```go
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
```
* `createMovie` function is used for creating a new movie and returning using json marshalling.
```go
// create movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
```
* `updateMovie` function is used for updating a particular movie and returning using json marshalling.
```go
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
```

## main function
This is the entry point of the go program.

```go
func main() {
	r := mux.NewRouter() // creates a new router

	// lets add some movies
	movies = append(movies, Movie{ID: strconv.Itoa(rand.Intn(100000000)), Isbn: strconv.Itoa(rand.Intn(6969)), Title: "SVSC", Director: &Director{Firstname: "Srikanth", Lastname: "Addala"}})
	movies = append(movies, Movie{ID: strconv.Itoa(rand.Intn(100000000)), Isbn: strconv.Itoa(rand.Intn(6969)), Title: "hey pillagada", Director: &Director{Firstname: "Sameer", Lastname: "Thahir"}})
	movies = append(movies, Movie{ID: strconv.Itoa(rand.Intn(100000000)), Isbn: strconv.Itoa(rand.Intn(6969)), Title: "orange", Director: &Director{Firstname: "Bommarillu", Lastname: "Bhaskar"}})

    // movie routes handling various functions
	r.HandleFunc("/movies", getMovies).Methods("GET")           // get all movies
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // get a specific movie
	r.HandleFunc("/movies", createMovie).Methods("POST")        // create a new movie
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // update an existing movie
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // delete a movie

	// golang server
	fmt.Printf("Starting golang server on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
```

## social links 
[![LinkedIn](https://img.shields.io/badge/LinkedIn-connect-0077B5?style=for-the-badge&logo=linkedin)](https://www.linkedin.com/in/villuri-jayavamsi) 
[![GitHub](https://img.shields.io/badge/GitHub-v4r5v4m5-181717?style=for-the-badge&logo=github)](https://github.com/v4r5v4m5) 
[![Twitter](https://img.shields.io/badge/Twitter-v-1DA1F2?style=for-the-badge&logo=twitter)](https://twitter.com/v1llur1) 