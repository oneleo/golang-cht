package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

func main() {
	h := &app.Handler{
		Title:  "Hello Demo",
		Author: "Maxence Charriere",
	}

	if err := http.ListenAndServe(":7777", h); err != nil {
		panic(err)
	}
}
