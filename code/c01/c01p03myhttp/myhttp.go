// +build OMIT
package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler struct{}

// START OMIT
func main() {
	p := ":8080"
	http.HandleFunc("/", sayHello) // HL
	log.Printf("Starting server...!\nPlease open: http://127.0.0.1%s", p)
	log.Fatal(http.ListenAndServe(p, nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) { // HL
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}

// END OMIT
