package zerolog

import (
	"io"

	"github.com/rs/zerolog"
)

// NewLogger creates a new zerolog logger with an optional pretty
// print format.
func NewLogger(output io.Writer, prettyPrint bool) zerolog.Logger {
	if prettyPrint {
		output = zerolog.ConsoleWriter{Out: output}
	}
	return zerolog.New(output).With().
		Timestamp().
		Str("service", "api").
		Logger()
}
