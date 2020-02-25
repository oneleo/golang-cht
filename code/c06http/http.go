// +build OMIT
package main

import (
	"log"
	"net/http"
)

type myHandler struct{}

func main() {
	httpServer()
}

func httpServer() { // OMIT START
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", sayHello) // HL
	mux.Handle("/", &myHandler{})      // HL

	mux.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Starting server...\nPlease open: http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func sayHello(w http.ResponseWriter, r *http.Request) { // HL
	w.Write([]byte("Hello world!!"))
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { // HL
	w.Write([]byte("The request URL is: " + r.URL.String()))
	log.Println(r.URL.String() + " was triggered")
} // OMIT END
