package exception

import "errors"

var (
	// ErrorTextNotFound is an error representing request not found.
	ErrorTextNotFound = errors.New("api.msg.error.common.not_found")

	// ErrorTextInternalServerError is an error representing internal server error.
	ErrorTextInternalServerError = errors.New("api.msg.error.common.internal_server_error")
)
