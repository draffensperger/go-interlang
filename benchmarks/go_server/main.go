package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func harmonic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	iStr := path[1:len(path)]
	if iStr == "stop" {
		os.Exit(0)
	}
	i, err := strconv.Atoi(iStr)
	if err == nil {
		fmt.Fprintf(w, "%v", 1.0/float64(i))
	}
}

func main() {
	http.HandleFunc("/", harmonic)
	http.ListenAndServe(":8000", nil)
}
