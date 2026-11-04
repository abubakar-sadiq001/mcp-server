package jsonrpc

const (
	ErrorParseError     = -32700 // Invalid JSON
	ErrorInvalidRequest = -32600 // Invalid Request onject
	ErrorMethodNotFound = -32602 // Method does not exist
	ErrorInvalidParams  = -32602 // Invalid method parameters
	ErrorInternal       = -32603 // Interbal JSON-RPC error
)
