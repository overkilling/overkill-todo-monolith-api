package http

import "net/http"

// Recoverer middleware for capturing panics and delegating to
// appropriate obsevability handlers. Returns a HTTP 500 status code.
func Recoverer() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if panicReason := recover(); panicReason != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
