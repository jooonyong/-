package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello form home!")
}
func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on https://localhost %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
