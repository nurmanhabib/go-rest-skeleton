package exception

import "errors"

var (
	// ErrorTextNotFound is an error representing request not found.
	ErrorTextNotFound = errors.New("api.msg.error.common.not_found")
)
