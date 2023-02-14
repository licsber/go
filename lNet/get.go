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

	if resp.StatusCode == http.StatusOK {
		return io.ReadAll(resp.Body)
	}

	if resp.StatusCode == http.StatusForbidden {
		return nil, ErrForbidden
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	if resp.StatusCode == http.StatusBadGateway {
		return nil, ErrBadGateway
	}

	log.Println(resp)
	return nil, ErrUnknown
}
