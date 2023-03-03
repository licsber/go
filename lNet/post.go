package lNet

import (
	"bytes"
	"net/url"
)

func (c *Client) PostFormBytes(url string, data url.Values) ([]byte, error) {
	return c.Handler(c.PostForm(url, data))
}

func (c *Client) PostJSON(url string, b []byte) ([]byte, error) {
	return c.Handler(c.Post(url, "application/json", bytes.NewReader(b)))
}
