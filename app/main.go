package main

import (
	"fmt"
	"os"
)

func main() {
	p, exists := os.LookupEnv("PORT")

	if !exists {
		fmt.Println("The port has not been set.")
		os.Exit(1)
	}

	port := ":" + p

	startServer(port)
}
