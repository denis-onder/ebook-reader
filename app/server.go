package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func startServer(port string) {
	// Static files
	staticDir := http.Dir("../static")
	fs := http.FileServer(staticDir)
	http.Handle("/static/", fs)

	//Request handlers
	http.HandleFunc("/", booksHandler)

	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
