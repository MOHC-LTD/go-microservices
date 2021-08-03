package jwt

import "fmt"

// IncorrectAlgorithmError thrown when a jwt has the wrong algorithm.
type IncorrectAlgorithmError struct {
	msg string
}

func (e IncorrectAlgorithmError) Error() string {
	return e.msg
}

// NewIncorrectAlgorithmError builds a new IncorrectAlgorithmError
func NewIncorrectAlgorithmError(alg string) error {
	return IncorrectAlgorithmError{
		msg: fmt.Sprintf("Incorrect jwt algorithm: %v", alg),
	}
}

// IncorrectHeaderFormat thrown when an authorization header does not have the correct format.
type IncorrectHeaderFormat struct {
	msg string
}

func (e IncorrectHeaderFormat) Error() string {
	return e.msg
}

// NewIncorrectHeaderFormat builds a new IncorrectHeaderFormat
func NewIncorrectHeaderFormat() error {
	return IncorrectHeaderFormat{
		msg: "Authorization header has the wrong format, must be `Bearer <token>`",
	}
}

// FailedToDecodeClaims thrown when a jwt's claims cannot be decoded
type FailedToDecodeClaims struct {
	msg string
}

func (e FailedToDecodeClaims) Error() string {
	return e.msg
}

// NewFailedToDecodeClaims builds a new FailedToDecodeClaims
func NewFailedToDecodeClaims() error {
	return FailedToDecodeClaims{
		msg: "Failed to decode the claims for the jwt",
	}
}
