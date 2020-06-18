package http_test

import (
	realHttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/overkilling/overkill-todo-monolith-api/http"
	"github.com/stretchr/testify/assert"
)

var noErrorHandler = func(w realHttp.ResponseWriter, r *realHttp.Request) {
	w.WriteHeader(realHttp.StatusOK)
}

var panicHandler = func(realHttp.ResponseWriter, *realHttp.Request) {
	panic("some error")
}

func TestRecoverer(t *testing.T) {
	tt := []struct {
		name         string
		handler      realHttp.HandlerFunc
		expectedCode int
	}{
		{
			name:         "no error",
			handler:      noErrorHandler,
			expectedCode: realHttp.StatusOK,
		},
		{
			name:         "panic",
			handler:      panicHandler,
			expectedCode: realHttp.StatusInternalServerError,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := httptest.NewRecorder()

			middleware := http.Recoverer()
			middleware(realHttp.HandlerFunc(tc.handler)).ServeHTTP(res, nil)

			assert.Equal(t, tc.expectedCode, res.Code)
		})
	}
}
