package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
)

func HookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello, World!\n")
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, span := otel.Tracer("circleci").Start(context.Background(), "request")
	defer span.End()
	fmt.Fprintf(w, "Created Request\n")
}

func TraceparentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("received /traceparent request for '%v'", r.URL.Path)
}
func main() {
	http.HandleFunc("", HookHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
