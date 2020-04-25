package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func startServer() {
	p, exists := os.LookupEnv("PORT")

	if !exists {
		fmt.Println("The port has not been set.")
		os.Exit(1)
	}

	port := ":" + p

	dir := http.Dir("./ebooks")
	fs := http.FileServer(dir)
	http.Handle("/", fs)

	fmt.Println("Listening on port " + port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
