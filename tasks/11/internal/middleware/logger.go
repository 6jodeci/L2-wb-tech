package internal

import (
	"log"
	"net/http"
)

// Logger выводит логи в терминал
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %v %v %v\n", r.Method, r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
