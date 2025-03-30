package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	router := http.NewServeMux()
	
	server := http.Server{
		Addr:    ":8080",
		Handler: middleware(router),
	}

	fmt.Printf("Server is running on http://localhost%s\n", server.Addr)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.RequestURI, "/")
		name := "World"
		if len(parts) > 1 && parts[1] != "" {
			name = parts[1]
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	server.ListenAndServe()
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware")
		next.ServeHTTP(w, r)
	})
}