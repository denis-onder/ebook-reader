package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type book struct {
	Name string
}

type booksPage struct {
	Title         string
	Books         []book
	NumberOfBooks int
	Links         []string
}

func getBooks(dir string) ([]book, error) {
	var output []book

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		name := file.Name()
		runes := []rune(name)
		extension := string(runes[len(runes)-4:])
		if extension != ".pdf" {
			continue
		}
		output = append(output, book{Name: string(runes[:len(runes)-4])})
	}

	return output, nil
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := getBooks("../static/books")

	if err != nil {
		fmt.Fprintf(w, "<h1>No books found.</h1>")
	}

	p := booksPage{Title: "Ebook Reader", Books: books, NumberOfBooks: len(books)}

	t, err := template.ParseFiles("../static/html/index.html")

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, p)
}
