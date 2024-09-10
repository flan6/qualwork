package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"strings"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tp := template.Must(template.ParseFiles("assets/index.html"))

		if err := tp.Execute(w, nil); err != nil {
			slog.Error(fmt.Errorf("failed to execute index.html: %w", err).Error())
		}
	})

	r.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tp := template.Must(template.ParseFiles("assets/contact.html"))

		err := tp.Execute(w, nil)
		if err != nil {
			slog.Error(fmt.Errorf("failed to execute contact.html: %w", err).Error())
		}
	})

	fs := http.FileServer(http.Dir("static"))

	r.Handle("/static/", http.StripPrefix("/static/", setContentTypeMiddleware(fs)))

	slog.Info("server listening on port 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error(fmt.Errorf("shuting down server: %w", err).Error())
	}
}

func setContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") {
			// Set Content-Type for CSS files
			w.Header().Set("Content-Type", "text/css")
		}
		// Serve the file
		next.ServeHTTP(w, r)
	})
}
