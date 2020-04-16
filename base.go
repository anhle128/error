package error

import "strconv"

// Message error
type Message struct {
	Message         string `json:"message,omitempty"`
	InternalMessage string `json:"internal_message,omitempty"`
	ErrorCode       string `json:"error_code,omitempty"`
}

// Error - return message
func (b Message) Error() string {
	return b.Message
}

// InternalError - return internal message
func (b Message) InternalError() string {
	return b.InternalMessage
}

// GetHTTPCode - return code
func (b Message) GetHTTPCode() int {
	httpCode, _ := strconv.Atoi(b.ErrorCode[:3])
	return httpCode
}

// GetErrorCode - return error code
func (b Message) GetErrorCode() string {
	return b.ErrorCode
}
