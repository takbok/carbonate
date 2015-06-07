package framework

import (
	"fmt"
	"net/http"

	"kodeclan.com/framework/utils"
)

type ControllerChild interface {
	SetParent(*BaseController)
}

type Controller interface {
	registerRequestAndResponse(*http.ResponseWriter, **http.Request)
	sendHttpResponseCode(int)
	respond(interface{})
}

type BaseController struct {
	module string
	index  string

	model *BaseModel
	view  *BaseView

	request  *http.Request
	response http.ResponseWriter
}

func (c *BaseController) SetModelAndView(m Model, v View) {
	m.SetParent(c)
	v.SetParent(c)

	v.setPackageName(utils.GetCallingPackageName())
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
