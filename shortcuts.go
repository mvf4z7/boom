// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package boom

import "net/http"

// BadRequest is shortcut for HTTP 400, Bad Request.
func BadRequest(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusBadRequest, format, msg...)
}

// Unauthorized is shortcut for HTTP 401, Unauthorized.
func Unauthorized(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusUnauthorized, format, msg...)
}

// Forbidden is shortcut for HTTP 403, Forbidden.
func Forbidden(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusForbidden, format, msg...)
}

// NotFound is shortcut for HTTP 404, Not Found.
func NotFound(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusNotFound, format, msg...)
}

// MethodNotAllowed is shortcut for HTTP 405, Method Not Allowed.
func MethodNotAllowed(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusMethodNotAllowed, format, msg...)
}

// NotAcceptable is shortcut for HTTP 406, Not Acceptable.
func NotAcceptable(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusNotAcceptable, format, msg...)
}

// ClientTimeout is shortcut for HTTP 408, Request Timeout.
func ClientTimeout(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusRequestTimeout, format, msg...)
}

// Conflict is shortcut for HTTP 409, Conflict.
func Conflict(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusConflict, format, msg...)
}

// Gone is shortcut for HTTP 410, Gone.
func Gone(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusGone, format, msg...)
}

// Error is shortcut for HTTP 500, Internal Server Error.
func Error(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusInternalServerError, format, msg...)
}

// NotImplemented is shortcut for HTTP 501, Not Implemented.
func NotImplemented(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusNotImplemented, format, msg...)
}

// BadGateway is shortcut for HTTP 502, Bad Gateway.
func BadGateway(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusBadGateway, format, msg...)
}

// Unavailable is shortcut for HTTP 503, Service Unavailable.
func Unavailable(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusServiceUnavailable, format, msg...)
}

// GatewayTimeout is shortcut for HTTP 504, Gateway Timeout.
func GatewayTimeout(format string, msg ...interface{}) *Boom {
	return CreateError(http.StatusGatewayTimeout, format, msg...)
}
