package framework

import (
	"fmt"
	"net/http"
)

type Controller interface {
	registerRequestAndResponse(*http.ResponseWriter, **http.Request)
	sendHttpResponseCode(int)
	respond(interface{})
}

type BaseController struct {
	model *BaseModel
	view  *BaseView

	request  *http.Request
	response http.ResponseWriter
}

func (c *BaseController) registerRequestAndResponse(w *http.ResponseWriter, r **http.Request) {
	c.request = *r
	c.response = *w
}

func (c *BaseController) sendHttpResponseCode(code int) {
	c.response.WriteHeader(code)
}

func (c *BaseController) respond(reply interface{}) {
	fmt.Fprintln(c.response, reply)
}
