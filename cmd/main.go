package main

import (
	"log"
	"net/http"

	"github.com/sohamsk/htmxPrime/internal/handlers"
	"github.com/sohamsk/htmxPrime/internal/middleware"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("POST /contacts", handlers.AddContact)

	http.ListenAndServe(":8080", middleware.RequestLoggerMiddleware(mux))
	log.Println("Server Started on port 8080")
}
