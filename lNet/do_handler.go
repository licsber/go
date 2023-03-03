package lNet

import (
	"io"
	"log"
	"net/http"
)

func DoHandler(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println("Fatal: close resp body failed.")
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

	log.Println("Unknown: response status code.")
	log.Println(resp)
	return nil, ErrUnknown
}
