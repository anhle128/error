package error

// Message error
type Message struct {
	Message         string `json:"message,omitempty"`
	InternalMessage string `json:"internal-message,omitempty"`
	Code            int    `json:"code,omitempty"`
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
	return b.Code
}
