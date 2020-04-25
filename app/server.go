package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func startServer(port string) {
	dir, exists := os.LookupEnv("BOOKS_DIR")

	if !exists {
		fmt.Println("The books directory has not been set.")
	}

	// Static files
	booksDir := http.Dir(dir)
	fs := http.FileServer(booksDir)
	http.Handle("/ebooks/", http.StripPrefix("/ebooks/", fs))

	//Request handlers
	http.HandleFunc("/", booksHandler)

	fmt.Println("Listening on port " + port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
