package zerolog

import (
	"net/http"

	"github.com/rs/zerolog/hlog"
)

// PanicHandler for logging recovered panics
type PanicHandler struct{}

// HandlePanic logs the panic reason as an error
func (handler *PanicHandler) HandlePanic(req *http.Request, panicReason interface{}) {
	logEvent := hlog.FromRequest(req).Error()
	if errorReason, ok := panicReason.(error); ok {
		logEvent.Stack().Err(errorReason)
	} else {
		logEvent.Interface("error", panicReason)
	}
	logEvent.Send()
}

// NewPanicHandler creates a handler with a given logger.
func NewPanicHandler() *PanicHandler {
	return &PanicHandler{}
}
