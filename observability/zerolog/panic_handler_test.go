package zerolog_test

import (
	"bytes"
	"errors"
	"fmt"
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
			panicHandler := zerolog.NewPanicHandler(log)

			panicHandler.HandlePanic(tc.panicReason)

			expectedLog := fmt.Sprintf(`{
				"level": "error",
				"error": "%s"
			}`, tc.expectedError)
			assert.JSONEq(t, expectedLog, string(out.Bytes()))
		})
	}
}
