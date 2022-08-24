package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users"))
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
