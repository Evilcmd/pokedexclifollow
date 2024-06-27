package pokeapi

import (
	"net/http"
	"time"
)

const BaseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		http.Client{
			Timeout: time.Minute,
		},
	}
}