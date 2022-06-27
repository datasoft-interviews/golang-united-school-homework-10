package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{name}", getName).Methods(http.MethodGet)
	router.HandleFunc("/bad", get500).Methods(http.MethodGet)
	router.HandleFunc("/data", getBody).Methods(http.MethodPost)
	router.HandleFunc("/headers", getSum).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func getName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Fprintf(w, "Hello, %s!", name)
}

func get500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func getBody(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "I got message:\n%s", body)
}

func getSum(w http.ResponseWriter, r *http.Request) {
	aS := r.Header.Get("a")
	bS := r.Header.Get("b")
	a, _ := strconv.Atoi(aS)
	b, _ := strconv.Atoi(bS)
	w.Header().Add("a+b", fmt.Sprint(a+b))
}
