package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "alive")
}

func rpnCalculator(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(time.Now(), err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	value, err := MakeStackCalc(string(body))
	if err != nil {
		log.Print(time.Now(), err)
		http.Error(w, "Bad request see https://en.wikipedia.org/wiki/Reverse_Polish_notation "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, value.Value())
}

func main() {
	port := 5000
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)
	r.HandleFunc("/calculate", rpnCalculator)
	fmt.Println("RPN server is up, Listening on Port ", port)
	log.Fatal(http.ListenAndServe("localhost:"+fmt.Sprint(port), r))
}
