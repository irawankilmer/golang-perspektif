package main

import (
	"fmt"
	"goperspektif/server"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.Index)

	fmt.Println("klik disini: http://localhost/8080")
	http.ListenAndServe(":8080", nil)
}
