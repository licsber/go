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

	switch resp.StatusCode {
	case http.StatusOK:
		return io.ReadAll(resp.Body)
	case http.StatusBadRequest:
		return nil, ErrBadRequest
	case http.StatusUnauthorized:
		return nil, ErrUnAuthorized
	case http.StatusForbidden:
		return nil, ErrForbidden
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusUnprocessableEntity:
		return nil, ErrUnprocessableEntity
	case http.StatusInternalServerError:
		return nil, ErrInternalServerError
	case http.StatusBadGateway:
		return nil, ErrBadGateway
	default:
		{
			log.Println("Unknown: response status code.")
			log.Println(resp)
			return nil, ErrUnknown
		}
	}
}
