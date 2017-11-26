// Copyright (C) 2015 Thomas de Zeeuw.
//
// Licensed under the MIT license that can be found in the LICENSE file.

package boom

import "fmt"

func ExampleBoom_Bytes() {
	notFoundErr := NotFound("My message")
	err := Error("My message")
	fmt.Println(string(notFoundErr.Bytes()))
	fmt.Println(string(err.Bytes()))
	// Output:
	// 404 Not Found: My message
	// 500 Internal Server Error: My message
}

func ExampleBoom_JSON() {
	notFoundErr := NotFound("My message")
	err := Error("My message")
	fmt.Println(string(notFoundErr.JSON()))
	fmt.Println(string(err.JSON()))
	// Output:
	// {"statusCode":404,"error":"Not Found","message":"My message"}
	// {"statusCode":500,"error":"Internal Server Error","message":"My message"}
}

func ExampleBoom_XML() {
	notFoundErr := NotFound("My message")
	err := Error("My message")
	fmt.Println(string(notFoundErr.XML()))
	fmt.Println(string(err.XML()))
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?><error><status_code>404</status_code><error>Not Found</error><message>My message</message></error>
	// <?xml version="1.0" encoding="UTF-8"?><error><status_code>500</status_code><error>Internal Server Error</error><message>My message</message></error>
}
