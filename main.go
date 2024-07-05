package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"text/template"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tp := template.Must(template.ParseFiles("assets/index.html"))

		if err := tp.Execute(w, nil); err != nil {
			slog.Error(fmt.Errorf("failed to execute index.html: %w", err).Error())
		}
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 3)

		tp := template.Must(template.ParseFiles("assets/hello.html"))

		err := tp.Execute(w, map[string]string{
			"Message": "This is a dynamic message from Go!",
		})
		if err != nil {
			slog.Error(fmt.Errorf("failed to execute hello.html: %w", err).Error())
		}
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", setContentTypeMiddleware(fs)))

	slog.Info("server listening on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error(fmt.Errorf("failed to init server: %w", err).Error())
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
