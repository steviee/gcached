package main

import (
	"log"
	"net/http"
	"time"
)

// Logger wraps the HttpHandler for logging.
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := newLoggingResponseWriter(w)
		inner.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode

		log.Printf(
			"%s\t%d\t%s\t%s\t%s",
			r.Method,
			statusCode,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
