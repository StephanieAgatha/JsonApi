package main

import (
	"learn-clean-movie/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/lists", controllers.ShowAllList)
	http.HandleFunc("/movieid", controllers.FindMovieById)
	http.HandleFunc("/moviename", controllers.FindMovieByName)

	http.ListenAndServe("localhost:3000", nil)
}
