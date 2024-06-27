package services

import (
	"encoding/json"
	"net/http"
)

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func Movies() []Movie {
	movs := []Movie{
		{1, "Spiderman", 2002},
		{2, "John Wick", 2004},
		{3, "The Raid", 2011},
		{4, "The Raid 2", 2014},
	}

	return movs
}
func GetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		movies := Movies()
		dataMovies, err := json.Marshal(movies)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMovies)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}
