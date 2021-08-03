package keys

import "fmt"

// Store is a place to store keys
type Store interface {
	// Set sets the value of a key
	Set(name string, key string) error
	// Get gets the a key by its name
	Get(name string) (string, error)
}

// KeyNotFound is thrown when a key can't be found
type KeyNotFound struct {
	name string
}

func (err KeyNotFound) Error() string {
	return fmt.Sprintf("No key with name %v", err.name)
}

// NewKeyNotFound builds the `KeyNotFound` error
func NewKeyNotFound(name string) error {
	return KeyNotFound{
		name,
	}
}
