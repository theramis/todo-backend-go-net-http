package main

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func addAllMiddlewares(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return handlers.RecoveryHandler()(
		handlers.LoggingHandler(os.Stdout,
			contentTypeMiddleware(
				corsMiddleware(
					alwaysRespondToOptionsMiddleware(
						http.HandlerFunc(next))))))
}

func contentTypeMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

func corsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("access-control-allow-headers", "accept, content-type")
		next.ServeHTTP(w, r)
	}
}

func alwaysRespondToOptionsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
