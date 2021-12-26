package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page!")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is user page!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is user creation page!")
}

func order(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is order page!")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", user)
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/order", order)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
