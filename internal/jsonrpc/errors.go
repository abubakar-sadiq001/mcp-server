package jsonrpc

const (
	ErrorParseError     = -32700 // Invalid JSON
	ErrorInvalidRequest = -32600 // Invalid Request onject
	ErrorMethodNotFound = -32602 // Method does not exist
	ErrorInvalidParams  = -32602 // Invalid method parameters
	ErrorInternal       = -32603 // Interbal JSON-RPC error
)

// NewError creates a new JSON-RPC error
func NewError(code int, message string, data any) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewParseError creates a parse error
func NewParseError(data any) *Error {
	return NewError(ErrorParseError, "Parse error", data)
}

func NewMethodNotFoundError(method string) *Error {
	return NewError(ErrorMethodNotFound, "Method not found error"+method, nil)
}

func NewInvalidParamsError(message string) *Error {
	return NewError(ErrorInvalidParams, message, nil)
}

func NewInternalError(err error) *Error {
	return NewError(ErrorInternal, "Internal error", err.Error())
}
