package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/family", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/healthcheck", HealthCheck)
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
func Secret(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1> Hello Full Cycle</h1>"))
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "User:  %s Password %s", user, password)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("OK"))

	}
}
