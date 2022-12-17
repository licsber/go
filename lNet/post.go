package lNet

import (
	"bytes"
	"io"
	"log"
	"net/url"
)

func (c *Client) PostFormBytes(url string, data url.Values) ([]byte, error) {
	resp, err := c.PostForm(url, data)
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

func (c *Client) PostJSON(url string, b []byte) ([]byte, error) {
	resp, err := c.Post(url, "application/json", bytes.NewReader(b))
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
