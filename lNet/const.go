package lNet

import "errors"

var ErrForbidden = errors.New("lNet: 403 Forbidden")
var ErrNotFound = errors.New("lNet: 404 Not Found")
var ErrBadGateway = errors.New("lNet: 502 Bad Gateway")
var ErrUnknown = errors.New("lNet: Unknown error")
