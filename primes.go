package main

import (
	"fmt"
	"github.com/flower-pot/primes/calculation"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{from:[0-9]+}/{to:[0-9]+}", handler).Methods("GET")
	http.Handle("/", r)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	socket := fmt.Sprintf(":%v", port)
	log.Println(fmt.Sprintf("Listening on %v...", socket))
	log.Fatal(http.ListenAndServe(socket, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	from, from_err := strconv.Atoi(params["from"])
	to, to_err := strconv.Atoi(params["to"])
	if from_err != nil || to_err != nil {
		if from_err != nil {
			fmt.Fprintln(w, fmt.Sprintf("Error with from: %s", from_err))
		}
		if to_err != nil {
			fmt.Fprintln(w, fmt.Sprintf("Error with to: %s", to_err))
		}
	}

	primes, err := calculation.CalculatePrimes(from, to)
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("%s", err))
	}
	fmt.Fprintln(w, fmt.Sprintf("%v", primes))
}
