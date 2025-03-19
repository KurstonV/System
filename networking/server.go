package main

import (
	"log"
	"net/http"
)

func home(w http.Responsewriter, r http.Request){
	w.Write([]byte("Welcome to Home"))
}

func login(w http.Responsewriter, r http.Request){
	w.Write([]byte("Login"))
}

func password(w http.Responsewriter, r http.Request){
	w.Write([]byte("Enter the password"))
}

func main() {
	mux := http.NewServerMux()

	mux.HandlerFunc("/.home")
	mux.HandlerFunc("/login", login)
	mux.HandlerFunc("/password", password)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
