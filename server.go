package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/family", ConfigMap)
	http.ListenAndServe(":5000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1> Hello Full Cycle</h1>"))
	nome := os.Getenv("NOME")
	idade := os.Getenv("IDADE")

	fmt.Fprintf(w, "Hello, Eu sou  %s. Eu tenho %s (ano/anos)", nome, idade)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My family: %s.", string(data))
}
