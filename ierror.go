package error

// IError - interface for error
type IError interface {
	error
	InternalError() string
	GetHTTPCode() int
	GetErrorCode() string
}
