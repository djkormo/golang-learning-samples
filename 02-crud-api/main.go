package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math/rand"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies [] Movie


func getMovies (w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-type","application/json")
  json.NewEncoder(w).Encode(movies)
}

func getMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	params:= mux.Vars(r)
	for _, item:=range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
  }

func  deleteMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	params:= mux.Vars(r)
	for index,item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
}

func  createMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) 
	// create new movie ID
	movie.ID=strconv.Itoa(rand.Intn(1000000000))
	// Add new mowie
	movies= append(movies, movie)
	json.NewEncoder(w).Encode(movie)

	
}

func  updateMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	params:= mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	for index,item := range movies {
		if item.ID == params["id"] {
			// delete item with index number
			movies = append(movies[:index],movies[index+1:]...)
            movie.ID=item.ID
			// add updated movie
			movies = append (movies,movie)
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}



  

func main()  {
	fmt.Printf("Hello CRUD API\n")
    
	movies = append(movies, Movie {ID:"1", Isbn:"1111", Title:"Pulp Fiction",Director: &Director {Firstname:"Pulp",Lastname:"Fiction"} })
	movies = append(movies, Movie {ID:"2", Isbn:"2222", Title:"Stranger Things",Director: &Director {Firstname:"Stranger",Lastname:"Strangest"} })
	r := mux.NewRouter()
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	fmt.Printf("Starting web server at 8000 port\n")

	log.Fatal(http.ListenAndServe(":8000",r))

	 
}