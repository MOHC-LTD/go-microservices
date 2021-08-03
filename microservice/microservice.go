package microservice

import (
	"encoding/json"
	"fmt"
	"go-microservices/http"
	"go-microservices/keys"
)

// APIProxy represents a connection to an API proxy
type APIProxy interface {
	// RefreshKey refreshes the public key of the proxy
	RefreshKey() error
	// PublicKey returns the current public key of the proxy
	PublicKey() (string, error)
}

type apiProxy struct {
	createURL    func(endpoint string) string
	clientID     string
	clientSecret string
	keys         keys.Store
	client       http.Client
}

// NewAPIProxy creates a connection to an API proxy
func NewAPIProxy(url string, clientID string, clientSecret string) (APIProxy, error) {
	createURL := func(endpoint string) string {
		return fmt.Sprintf("%v/%v", url, endpoint)
	}

	proxy := apiProxy{
		createURL:    createURL,
		clientID:     clientID,
		clientSecret: clientSecret,
		keys:         keys.NewMemoryStore(make(map[string]string)),
		client: http.NewClient(http.RequestHeaders{
			{
				Name:  "Content-Type",
				Value: "application/json",
			},
		}),
	}

	err := proxy.RefreshKey()
	if err != nil {
		return apiProxy{}, err
	}

	return proxy, nil
}

func (proxy apiProxy) RefreshKey() error {
	session, err := proxy.createSession()
	if err != nil {
		return err
	}

	keyBytes, _, err := proxy.client.Get(proxy.createURL("public_key"), http.RequestHeaders{{
		Name:  "cookie",
		Value: session,
	}})
	if err != nil {
		return err
	}

	publicKey := string(keyBytes)

	proxy.keys.Set("publicKey", publicKey)

	return nil
}

func (proxy apiProxy) PublicKey() (string, error) {
	key, err := proxy.keys.Get("publicKey")
	if err != nil {
		return "", err
	}

	return key, nil
}

func (proxy apiProxy) createSession() (string, error) {
	body := struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{
		ClientID:     proxy.clientID,
		ClientSecret: proxy.clientSecret,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	_, headers, err := proxy.client.Post(proxy.createURL("auth/signin"), bodyBytes, nil)
	if err != nil {
		return "", err
	}

	session := headers.Get("Set-Cookie")

	return session, err
}
