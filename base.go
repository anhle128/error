package error

import "strconv"

// Message error
type Message struct {
	Message         string `json:"message,omitempty"`
	InternalMessage string `json:"internal-message,omitempty"`
	ErrorCode       string `json:"code,omitempty"`
}

// Error - return message
func (b Message) Error() string {
	return b.Message
}

// InternalError - return internal message
func (b Message) InternalError() string {
	return b.InternalMessage
}

// HTTPCode - return code
func (b Message) HTTPCode() int {
	httpCode, _ := strconv.Atoi(b.ErrorCode[:3])
	return httpCode
}
