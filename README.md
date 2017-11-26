# Boom

[![GoDoc](https://godoc.org/github.com/Thomasdezeeuw/boom?status.svg)](https://godoc.org/github.com/Thomasdezeeuw/boom)

> inspired by [Nodejs Boom](https://github.com/hapijs/boom)

Boom creates easy to use JSON HTTP errors for use in api's, an example:

```go
package main

import (
	"os"

	"github.com/Thomasdezeeuw/boom"
)

func main() {
	// Creates a struct with:
	// StatusCode int
	// Error      string
	// Message    string
	notFoundError := boom.NotFound("The page you're looking for ain't here")

	// Or with a formatted message
	notFoundError = boom.NotFound("The %s you're looking for ain't here", "document")

	// Then write the message in JSON, or XML, or plain text.
	json := notFoundError.JSON()
	os.Stdout.Write(json)
}

```
