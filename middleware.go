package main

import (
	"net/http"
)

var ALLOWED_ORIGINS = []string{"http://localhost:5173", "https://go-pong-react-client-xsal4.ondigitalocean.app"}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for _, o := range ALLOWED_ORIGINS {
			if r.Header.Get("Origin") == o {
				w.Header().Set("Access-Control-Allow-Origin", o)
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Expose-Headers", "Link")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)

	})

}
