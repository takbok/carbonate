package brahma

import (
	"net/http"
	"reflect"

	"bitbucket.org/takbok/brahma/router"
)

// Dispatch the request to the designated module and method.
// This function is invoked for each request
func DispatchRequestViaURL(url string, c Controller, w *http.ResponseWriter, r **http.Request) {
	// Register the request and response object
	c.registerRequestAndResponse(w, r)

	// Get the string response for the current reqest
	response := dispatchRequestAndGetResponse(c, url)

	// Write the response to the default response channel
	c.respond(response)
}

// Dispatch a request to it's controller and fetch the response
func dispatchRequestAndGetResponse(c Controller, url string) string {
	// Get the request parameters, the method name and the arguments
	_, function, args := router.GetParamsFromURL(url)

	// Get the controller and it's requested method
	controller := reflect.ValueOf(c)
	method := controller.MethodByName(function)

	if !method.IsValid() {
		// HTTP 404 if the method isn't found
		c.sendHttpResponseCode(http.StatusNotFound)
		return "This page cannot be found"
	}

	mType := method.Type()
	if mType.NumOut() < 0 || mType.Out(0).String() != "string" {
		// HTTP 405 if the method called isn't allowed
		c.sendHttpResponseCode(http.StatusMethodNotAllowed)
		return "This page is not accessible"
	}

	if !method.Type().IsVariadic() && method.Type().NumIn() != len(args) {
		// HTTP 400 if URL is incomplete
		c.sendHttpResponseCode(http.StatusBadRequest)
		return "Illegal request"
	}

	in := make([]reflect.Value, len(args))

	for k, v := range args {
		in[k] = reflect.ValueOf(v)
	}

	// Call the method with the arguments and return the response as string
	response := method.Call(in)[0]
	return response.String()
}
