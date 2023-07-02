package controllers

import (
	"encoding/json"
	"learn-clean-movie/entities"
	"net/http"
	"strings"
)

var data = entities.MovieData

func ShowAllList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	if r.Method == "GET" {
		jsonresp, err := json.Marshal(data)
		if err != nil {
			panic(err)
			return
		}

		w.Write(jsonresp)
	}
}

func FindMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	if r.Method == "GET" {
		id := r.FormValue("id")
		for _, each := range data {
			if each.Id == id {
				jsonresp, err := json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				w.Write(jsonresp)
				return
			}
		}
		w.WriteHeader(404)
		w.Write([]byte("Movie ID Not Found"))
		return
	}
	w.WriteHeader(500)
	w.Write([]byte("Method Not Allowed >:("))
	return
}

func FindMovieByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	if r.Method == "GET" {
		name := r.FormValue("name")
		for _, each := range data {
			if strings.Contains(strings.ToLower(each.Name), strings.ToLower(name)) {
				jsonresp, err := json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				w.Write(jsonresp)
				return
			}
		}
		w.WriteHeader(404)
		w.Write([]byte("Movie Not Found"))
		return
	}
	w.WriteHeader(500)
	w.Write([]byte("Method Not Allowed"))
	return
}
