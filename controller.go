package brahma

import (
	"fmt"
	"net/http"

	"bitbucket.org/takbok/brahma/utils"
)

// Define the ControllerChild interface
type ControllerChild interface {
	setParent(*BaseController)
}

// Define the controller interface
type Controller interface {
	registerRequestAndResponse(*http.ResponseWriter, **http.Request)
	sendHttpResponseCode(int)
	respond(interface{})
}

// The base controller structure. It contains:
// - the address of the base model
// - the address of the base view
// - the address of the original request
// - the response writer object
type BaseController struct {
	model *BaseModel
	view  *BaseView

	request  *http.Request
	response http.ResponseWriter
}

// Set the base model and view for the controller
func (c *BaseController) SetModelAndView(m Model, v View) {
	// A rather obtuse way of getting things right
	// However, this has been thought through, and it is how you get things
	// working correctly in Go
	m.setParent(c)
	v.setParent(c)

	// Set the module name for the view
	v.setPackageName(utils.GetCallingPackageName())
}

// Register the request and response object
func (c *BaseController) registerRequestAndResponse(w *http.ResponseWriter, r **http.Request) {
	c.request = *r
	c.response = *w
}

// Send a HTTP status code with the default code explanation
func (c *BaseController) sendHttpResponseCode(code int) {
	c.response.WriteHeader(code)
}

// Send a reply back to the default response channel
func (c *BaseController) respond(reply interface{}) {
	fmt.Fprintln(c.response, reply)
}
