package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"sync"

	"go.opentelemetry.io/otel/trace"
)

type hookHandler struct {
	mu     sync.Mutex
	tracer trace.Tracer
}

func (h *hookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	_, span := h.tracer.Start(context.Background(), "request")
	defer span.End()
	defer fmt.Fprintf(w, "Created Request\n")
}

func main() {
	http.Handle("", new(hookHandler))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
