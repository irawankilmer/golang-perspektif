package server

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("nama")
	fmt.Fprintln(w, "Ini adalah halaman home", data)
}
