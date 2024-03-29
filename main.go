package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Lap Time Leaderboard!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}