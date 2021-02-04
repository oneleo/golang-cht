// +build OMIT
package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
func main() {
	p, m := ":8080", http.NewServeMux()
	m.HandleFunc("/", helloHandler) // HL
	log.Printf("Starting server...\nPlease open: http://127.0.0.1%s", p)
	log.Fatal(http.ListenAndServe(p, m))
}

func helloHandler(w http.ResponseWriter, r *http.Request) { // HL
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}

// END OMIT
