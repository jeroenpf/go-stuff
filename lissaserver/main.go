// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"learn/lissaserver/utilities"
	"log"
	"net/http"
	"strconv"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {

		cycles, _ := strconv.Atoi(r.URL.Query().Get("cycles"))
		if cycles > 100 || cycles <= 0 {
			cycles = 5
		}

		utilities.Lissajous(w, float64(cycles))
	}

	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
