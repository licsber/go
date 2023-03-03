package lNet

import "net/http"

type Client struct {
	http.Client
	Handler func(*http.Response, error) ([]byte, error)
}

func NewClient() *Client {
	return &Client{
		Handler: DoHandler,
	}
}
