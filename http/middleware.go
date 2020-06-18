package http

import "net/http"

// RecovererPanicHandler is registered in the Recoverer middleware and
// is called whenever the middleware recovers from a panic.
// Useful to log an error message or to report an error metric.
type RecovererPanicHandler interface {
	HandlePanic(req *http.Request, panicReason interface{})
}

// Recoverer middleware for capturing panics and delegating to
// appropriate obsevability handlers. Returns a HTTP 500 status code.
func Recoverer(panicHandlers ...RecovererPanicHandler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				if panicReason := recover(); panicReason != nil {
					for _, handler := range panicHandlers {
						handler.HandlePanic(req, panicReason)
					}

					w.WriteHeader(http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, req)
		})
	}
}
