package zerolog_test

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/overkilling/overkill-todo-monolith-api/observability/zerolog"
	realZerolog "github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	tt := []struct {
		name          string
		panicReason   interface{}
		expectedError string
	}{
		{
			name:          "string reason",
			panicReason:   "some string reason",
			expectedError: "some string reason",
		},
		{
			name:          "err reason",
			panicReason:   errors.New("some err reason"),
			expectedError: "some err reason",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			log := realZerolog.New(out).With().Logger()
			req, _ := http.NewRequest("GET", "http://localhost:3000/some-endpoint", nil)
			req = req.WithContext(log.WithContext(req.Context()))

			zerolog.NewPanicHandler().HandlePanic(req, tc.panicReason)

			expectedLog := fmt.Sprintf(`{
				"level": "error",
				"error": "%s"
			}`, tc.expectedError)
			assert.JSONEq(t, expectedLog, string(out.Bytes()))
		})
	}
}
