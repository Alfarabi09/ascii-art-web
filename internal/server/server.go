package server

import (
	"asciiart/internal/handler"
	"fmt"
	"net/http"
)

func Server() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.MainPage)
	mux.HandleFunc("/ascii-art", handler.PostArt)

	// mux.Handle("/template/css/", http.StripPrefix("/template/css/", http.FileServer(http.Dir("./template/css/"))))
	mux.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template/"))))
	fmt.Println("Server listening...")
	fmt.Println("http://localhost:7777")
	return http.ListenAndServe(":7777", mux)
}
