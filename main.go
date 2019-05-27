package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/davidharting/coordinates/coordinates"
	"github.com/gorilla/mux"
)

func getPolarCoordinates(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	x, errX := strconv.ParseFloat(v.Get("x"), 64)
	y, errY := strconv.ParseFloat(v.Get("y"), 64)

	if errX != nil || errY != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Must provide query params X and Y, which must be numbers.")
		return
	}

	cartesian := coordinates.CartesianPair{X: x, Y: y}
	polar := cartesian.Convert()
	response, err := json.Marshal(polar)

	if err != nil {
		w.WriteHeader(500)
		fmt.Print("Server error")
	}

	fmt.Fprint(w, string(response))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/polar-coordinates", getPolarCoordinates).Methods("GET").Name("polar-coordinates")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
