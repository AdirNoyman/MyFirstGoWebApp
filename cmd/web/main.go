package main

import (
	"fmt"
	"hello_world3/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	// Routes
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// Server
	fmt.Println(fmt.Sprintf("Starting application on port %s 😎🤟", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
