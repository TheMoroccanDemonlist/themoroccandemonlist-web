package middleware

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%s %s %s", request.Method, request.URL.Path, request.RemoteAddr)
		next.ServeHTTP(writer, request)
	})
}
