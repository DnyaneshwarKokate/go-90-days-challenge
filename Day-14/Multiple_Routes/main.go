package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Page")
}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Println("Server running on 8080")
	http.ListenAndServe(":8080", nil)
}