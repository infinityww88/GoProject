package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ret := r.URL.Query()
	a, _ := strconv.ParseInt(ret.Get("a"), 10, 32)
	b, _ := strconv.ParseInt(ret.Get("b"), 10, 32)
	fmt.Fprintf(w, "URL.Path = %q, ret = %d\n",
		r.URL.Path, a+b)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
