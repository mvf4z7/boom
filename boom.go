// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

// Package boom creates easy to use JSON HTTP errors for use in api's.
package boom

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

// Boom type with the HTTP Status Code, HTTP error message and
// a user given error message.
type Boom struct {
	XMLName    string `json:"-" xml:"error,-"`
	StatusCode int    `json:"statusCode" xml:"status_code"`
	Error      string `json:"error" xml:"error"`
	Message    string `json:"message" xml:"message"`
}

// Bytes return Boom in plain text. The following format is used:
// 	StatusCode Error: Message
func (b *Boom) Bytes() []byte {
	str := fmt.Sprintf("%d %s: %s", b.StatusCode, b.Error, b.Message)
	return []byte(str)
}

// JSON returns Boom in a JSON format.
func (b *Boom) JSON() []byte {
	buf, _ := json.Marshal(b)
	return buf
}

var xmlHeader = []byte(strings.TrimSpace(xml.Header))

// XML returns Boom in a XML format.
func (b *Boom) XML() []byte {
	buf, _ := xml.Marshal(b)
	return append(xmlHeader, buf...)
}

// CreateError create a formatted json error with the HTTP status code, HTTP error
// description and a provided user message.
//
// It's preferred to use some of the shortcuts provided.
func CreateError(statusCode int, format string, msg ...interface{}) *Boom {
	var message string

	if len(msg) == 0 {
		message = format
	} else {
		message = fmt.Sprintf(format, msg...)
	}

	return &Boom{
		StatusCode: statusCode,
		Error:      http.StatusText(statusCode),
		Message:    message,
	}
}
