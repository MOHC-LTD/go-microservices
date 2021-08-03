package keys

type memoryStore struct {
	keys map[string]string
}

// NewMemoryStore builds an in memory key store
func NewMemoryStore(initialKeys map[string]string) Store {
	var keys map[string]string

	if initialKeys != nil {
		keys = initialKeys
	} else {
		keys = make(map[string]string)
	}

	return memoryStore{
		keys,
	}
}

func (store memoryStore) Set(name string, key string) error {
	store.keys[name] = key

	return nil
}

func (store memoryStore) Get(name string) (string, error) {
	key := store.keys[name]

	if key == "" {
		return "", NewKeyNotFound(name)
	}

	return key, nil
}
