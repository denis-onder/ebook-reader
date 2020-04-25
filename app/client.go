package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type book struct {
	name string
}

type booksPage struct {
	books []book
}

func getBooks(dir string) ([]book, error) {
	var output []book

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		output = append(output, book{name: file.Name()})
	}

	return output, nil
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	dir, exists := os.LookupEnv("BOOKS_DIR")

	if !exists {
		fmt.Fprintf(w, "The books directory has not been set.")
	}

	books, err := getBooks(dir)

	if err != nil {
		fmt.Fprintf(w, "<h1>No books found.</h1>")
	}

	for i, book := range books {
		fmt.Fprintf(w, "<h1>%d. %s</h1>", i+1, book.name)
	}
}
