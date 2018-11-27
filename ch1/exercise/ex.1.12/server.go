package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ahetmanski/gobook/ch1/lissajous"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var cycles float64 = 5
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	if _, exists := r.Form["cycles"]; exists {
		if val, err := strconv.ParseFloat(r.Form["cycles"][0], 64); err != nil {
			log.Print(err)
		} else {
			cycles = val
		}
	}
	lissajous.Lissajous(w, cycles)
}
