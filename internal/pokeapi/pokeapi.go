package pokeapi

import (
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

// Client to make a http request in go, a http client object need to be present in memory, default client does not have
// a timeout, hence using a custom client here
type Client struct {
	httpClient http.Client
}

// NewClient initialize custom http client with a set timeout value
func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
