package zerolog

import "github.com/rs/zerolog"

// PanicHandler for logging recovered panics
type PanicHandler struct {
	log zerolog.Logger
}

// HandlePanic logs the panic reason as an error
func (handler *PanicHandler) HandlePanic(panicReason interface{}) {
	logEvent := handler.log.Error()
	if errorReason, ok := panicReason.(error); ok {
		logEvent.Stack().Err(errorReason)
	} else {
		logEvent.Interface("error", panicReason)
	}
	logEvent.Send()
}

// NewPanicHandler creates a handler with a given logger.
func NewPanicHandler(log zerolog.Logger) *PanicHandler {
	return &PanicHandler{log}
}
