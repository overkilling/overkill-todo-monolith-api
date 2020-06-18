package zerolog_test

import (
	"bytes"
	"testing"

	"github.com/overkilling/overkill-todo-monolith-api/observability/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	tt := []struct {
		name               string
		prettyPrint        bool
		expectedLogSnippet string
	}{
		{
			name:               "pretty log",
			prettyPrint:        true,
			expectedLogSnippet: "\x1b[0m \x1b[32mINF\x1b[0m test log \x1b[36mservice=\x1b[0mapi \x1b[36msome_option=\x1b[0msome_value\n",
		},
		{
			name:               "json log",
			prettyPrint:        false,
			expectedLogSnippet: `{"level":"info","service":"api","some_option":"some_value",`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			log := zerolog.NewLogger(out, tc.prettyPrint)

			log.Info().
				Str("some_option", "some_value").
				Msg("test log")

			assert.Contains(t, string(out.Bytes()), tc.expectedLogSnippet)
		})
	}
}
