package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":5000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1> Hello Full Cycle</h1>"))
	nome  := os.Getenv("NOME")
	idade := os.Getenv("IDADE")

	fmt.Fprintf(w, "Hello, Eu sou  %s. Eu tenho %s (ano/anos)", nome, idade)
}
