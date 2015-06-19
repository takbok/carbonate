// The Takbok Module package contains methods to set various properties for each module
package module

import (
	"net/http"

	"kodeclan.com/framework/router"
	"kodeclan.com/framework/utils"
)

// Set the default controller method for this module
func DefaultMethod(method string) {
	module := utils.GetCallingPackageName()
	router.DefaultMethodForModule(module, method)
}

// Register the request handler for the module.
// All requests to the module will be handled by this function
func RegisterRequestHandler(handler http.HandlerFunc) {
	module := utils.GetCallingPackageName()
	router.AddModuleHandler(module, handler)
}

// Add a static route to a controller method
func ForwardPermalinkToMethod(url, method string) {
	module := utils.GetCallingPackageName()
	router.AddStaticRoute(url, module, method)
}

// Set the invoker module as the default module
// All request to "/" will be handled by this module
func SetAsDefault() {
	module := utils.GetCallingPackageName()
	router.SetDefaultModule(module)
}
