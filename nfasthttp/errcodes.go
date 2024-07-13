package nfasthttp

import "errors"

// Clients replies
var ErrClientBadRequest = errors.New("bad request")         // Final status
var ErrClientUnavailable = errors.New("client unavailable") // Temporary status -> repeat routing
