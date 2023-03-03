package lNet

import (
	"net/http"
	"net/url"
)

func (c *Client) GetBytes(url string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	return c.Handler(c.Do(req))
}
