package error

// IError - interface for error
type IError interface {
	error
	InternalError() string
	HTTPCode() int
	GetErrorCode() string
}
