package lNet

import (
	"io"
	"log"
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

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	return io.ReadAll(resp.Body)
}
