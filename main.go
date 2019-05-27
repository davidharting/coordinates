package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getPolarCoordinates(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	x, errX := strconv.ParseFloat(v.Get("x"), 64)
	y, errY := strconv.ParseFloat(v.Get("y"), 64)

	if errX != nil || errY != nil {
		fmt.Fprintf(w, "Must provide query params X and Y, which must be numbers.")
	} else {
		fmt.Fprintf(w, "Okay then you sent us (%v, %v)", x, y)
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/polar-coordinates", getPolarCoordinates).Methods("GET").Name("polar-coordinates")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
