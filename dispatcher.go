package framework

import (
	"net/http"
	"reflect"

	"kodeclan.com/framework/router"
)

func DispatchRequestViaURL(url string, c Controller, w *http.ResponseWriter, r **http.Request) {
	c.registerRequestAndResponse(w, r)
	response := dispatchRequestAndGetResponse(c, url)

	c.respond(response)
}

func dispatchRequestAndGetResponse(c Controller, url string) string {
	_, function, args := router.GetParamsFromURL(url)

	controller := reflect.ValueOf(c)
	method := controller.MethodByName(function)

	if !method.IsValid() {
		c.sendHttpResponseCode(http.StatusNotFound)
		return "This page cannot be found"
	}

	mType := method.Type()
	if mType.NumOut() < 0 || mType.Out(0).String() != "string" {
		c.sendHttpResponseCode(http.StatusMethodNotAllowed)
		return "This page is not accessible"
	}

	if !method.Type().IsVariadic() && method.Type().NumIn() != len(args) {
		c.sendHttpResponseCode(http.StatusBadRequest)
		return "Illegal request"
	}

	in := make([]reflect.Value, len(args))

	for k, v := range args {
		in[k] = reflect.ValueOf(v)
	}

	response := method.Call(in)[0]
	return response.String()
}
