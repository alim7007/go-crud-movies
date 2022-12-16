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

type Movie struct{
ID string `json:"id"`
Isbn string `json:"isbn"`
Title string `json:"title"`
Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	}

var movies []Movie

func getMovies(resp http.ResponseWriter, req *http.Request){
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(movies)
}

func deleteMovie(resp http.ResponseWriter, req *http.Request){
	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(resp).Encode(movies)
}

func getMovie(resp http.ResponseWriter, req *http.Request){
	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(resp).Encode(item)
			return
		}
	}
}

func createMovie(resp http.ResponseWriter, req * http.Request){
	resp.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(resp).Encode(movie)

}

func updateMovie(resp http.ResponseWriter, req * http.Request){
	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(resp).Encode(item)
			return
		}
	}
}

func main(){
	r:= mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn:"194108", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID:"2", Isbn:"194110", Title: "Movie two", Director: &Director{Firstname: "Chris", Lastname: "Ll"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}