package errors

// The purpose of this file is to return sane error messages to users, rather than
// whatever message comes from the library functions that throw the original error.

const (
	decodeError              = "failed to decode the response from the GM API"
	jsonMarshalError         = "failed to marshal the request body JSON"
	postError                = "POST request to GM API failed"
	unidentifiedCommandError = "could not identify engine command"
)

// DecodeError is a wrapper for JSON decoding errors
type DecodeError struct {
	message string
}

// NewDecodeError creates a new DecodeError
func NewDecodeError() *DecodeError {
	return &DecodeError{
		message: decodeError,
	}
}

// Error returns the decode error message
func (e *DecodeError) Error() string {
	return e.message
}

// JSONMarshalError is a wrapper for JSON marshaling errors
type JSONMarshalError struct {
	message string
}

// NewJSONMarshalError creates a new JSONMarshalError
func NewJSONMarshalError() *JSONMarshalError {
	return &JSONMarshalError{
		message: jsonMarshalError,
	}
}

// Error returns the JSON marshal error message
func (e *JSONMarshalError) Error() string {
	return e.message
}

// PostError is a wrapper for POST request errors
type PostError struct {
	message string
}

// NewPostError creates a new PostError
func NewPostError() *PostError {
	return &PostError{
		message: postError,
	}
}

// Error returns the POST error message
func (e *PostError) Error() string {
	return e.message
}

// UnidentifiedCommandError is a wrapper for when an unidentified command is sent to the engine
type UnidentifiedCommandError struct {
	message string
}

// NewUnidentifiedCommandError creates a new UnidentifiedCommandError
func NewUnidentifiedCommandError() *UnidentifiedCommandError {
	return &UnidentifiedCommandError{
		message: unidentifiedCommandError,
	}
}

// Error returns the unidentified command error string
func (e *UnidentifiedCommandError) Error() string {
	return e.message
}
