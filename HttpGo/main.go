package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func serverIsHealthy() bool {
	// Check the health of the server and return ture of false accordingly
	// For example, check if the server can connect to the database
	return true
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Check the health of the server and return a status code accordingly
	if serverIsHealthy() {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Server is healthy\n")
		log.Printf("health check ok\n")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Server is not healthy")
		log.Printf("health check failed\n")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request " + r.URL.Host)
	ret := r.URL.Query()
	a, _ := strconv.ParseInt(ret.Get("a"), 10, 32)
	b, _ := strconv.ParseInt(ret.Get("b"), 10, 32)
	log.Printf("a = %d, b = %d\n", a, b)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "URL.Path = %q, ret = %d\n",
		r.URL.Path, a+b)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
