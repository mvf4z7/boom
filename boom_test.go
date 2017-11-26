// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package boom

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"testing"
)

type Test struct {
	statusCode int
	errMsg     string
	msgFormat  string
	msg        []interface{}
	fn         func(format string, msg ...interface{}) *Boom
}

func TestBoom(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{400, "Bad Request", "My message", []interface{}{}, BadRequest},
		{401, "Unauthorized", "My message", []interface{}{}, Unauthorized},
		{403, "Forbidden", "My message", []interface{}{}, Forbidden},
		{404, "Not Found", "My message", []interface{}{}, NotFound},
		{406, "Not Acceptable", "My message", []interface{}{}, NotAcceptable},
		{408, "Request Timeout", "My message", []interface{}{}, ClientTimeout},
		{409, "Conflict", "My message", []interface{}{}, Conflict},
		{410, "Gone", "My message", []interface{}{}, Gone},
		{500, "Internal Server Error", "My message", []interface{}{}, Error},
		{501, "Not Implemented", "My message", []interface{}{}, NotImplemented},
		{502, "Bad Gateway", "My message", []interface{}{}, BadGateway},
		{503, "Service Unavailable", "My message", []interface{}{}, Unavailable},
		{504, "Gateway Timeout", "My message", []interface{}{}, GatewayTimeout},
	}

	for _, test := range tests {
		if err := testBoom(test); err != nil {
			t.Fatal(err)
		}
	}
}

func TestFormattedMessages(t *testing.T) {
	t.Parallel()

	tests := []Test{
		{400, "Bad Request", "My %s message", []interface{}{"formatted"}, BadRequest},
		{401, "Unauthorized", "My %s message", []interface{}{"formatted"}, Unauthorized},
		{403, "Forbidden", "My %s message", []interface{}{"formatted"}, Forbidden},
		{404, "Not Found", "My %s message", []interface{}{"formatted"}, NotFound},
		{406, "Not Acceptable", "My %s message", []interface{}{"formatted"}, NotAcceptable},
		{408, "Request Timeout", "My %s message", []interface{}{"formatted"}, ClientTimeout},
		{409, "Conflict", "My %s message", []interface{}{"formatted"}, Conflict},
		{410, "Gone", "My %s message", []interface{}{"formatted"}, Gone},
		{500, "Internal Server Error", "My %s message", []interface{}{"formatted"}, Error},
		{501, "Not Implemented", "My %s message", []interface{}{"formatted"}, NotImplemented},
		{502, "Bad Gateway", "My %s message", []interface{}{"formatted"}, BadGateway},
		{503, "Service Unavailable", "My %s message", []interface{}{"formatted"}, Unavailable},
		{504, "Gateway Timeout", "My %s message", []interface{}{"formatted"}, GatewayTimeout},
	}

	for _, test := range tests {
		if err := testBoom(test); err != nil {
			t.Fatal(err)
		}
	}
}

func TestBytes(t *testing.T) {
	t.Parallel()

	var (
		statusCode = 404
		msg        = "my message"
		expected   = fmt.Sprintf("%d %s: %s", statusCode, http.StatusText(statusCode), msg)
	)

	b := CreateError(statusCode, msg)
	got := string(b.Bytes())

	if got != expected {
		t.Fatalf("Expected CreateError(%d, %s).Bytes() to return %s, but got %s",
			statusCode, msg, expected, got)
	}
}

func TestJSON(t *testing.T) {
	t.Parallel()

	var (
		statusCode = 404
		msg        = "my message"
		expected   = fmt.Sprintf(`{"statusCode":%d,"error":%q,"message":%q}`,
			statusCode, http.StatusText(statusCode), msg)
	)

	b := CreateError(statusCode, msg)
	got := string(b.JSON())

	if got != expected {
		t.Fatalf("Expected CreateError(%d, %s).JSON() to return %s, but got %s",
			statusCode, msg, expected, got)
	}
}

func TestXML(t *testing.T) {
	t.Parallel()

	var (
		statusCode = 404
		msg        = "my message"
		format     = `<?xml version="1.0" encoding="UTF-8"?><error>` +
			`<status_code>%d</status_code><error>%s</error>` +
			`<message>%s</message></error>`
		expected = fmt.Sprintf(format, statusCode, http.StatusText(statusCode), msg)
	)

	b := CreateError(statusCode, msg)
	got := string(b.XML())

	if got != expected {
		t.Fatalf("Expected CreateError(%d, %s).XML() to return %s, but got %s",
			statusCode, msg, expected, got)
	}
}

func testBoom(test Test) error {
	b := test.fn(test.msgFormat, test.msg...)

	if b.StatusCode != test.statusCode {
		fnName := getFnName(test.fn)
		return fmt.Errorf("Expected %s(\"%v\", %v).StatusCode to be %d, but got %d",
			fnName, test.msgFormat, test.msg, test.statusCode, b.StatusCode)
	}

	if b.Error != test.errMsg {
		fnName := getFnName(test.fn)
		return fmt.Errorf("Expected %s(\"%v\", %v).Error to be %s, but got %s",
			fnName, test.msgFormat, test.msg, test.errMsg, b.Error)
	}

	expectedMsg := fmt.Sprintf(test.msgFormat, test.msg...)
	if b.Message != expectedMsg {
		fnName := getFnName(test.fn)
		return fmt.Errorf("Expected %s(\"%v\", %v).Message to be %s, but got %s",
			fnName, test.msgFormat, test.msg, expectedMsg, b.Message)
	}

	return nil
}

func getFnName(fn interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	trim := "github.com/Thomasdezeeuw/boom."
	return fullName[len(trim):]
}
