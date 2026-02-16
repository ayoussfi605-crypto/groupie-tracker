package main

import (
	"fmt"
	"net/http"
	"groupie-tracker/handler"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/artist", handler.ArtistHandler)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
