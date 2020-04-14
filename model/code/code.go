package code

import "github.com/pkg/errors"

var (
	RequestErr        = errors.New("request error")
	UserIdNotFound    = errors.New("user id not found")
	SessionIdNotFound = errors.New("session id not found")
	MatchIdNotFound   = errors.New("match id not found")
)
