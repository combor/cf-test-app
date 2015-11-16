package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	variables := ""
	for _, e := range os.Environ() {
		variables = variables + e + "\n"
	}
	fmt.Fprintf(w, variables)
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
