package module

import (
	"net/http"

	"kodeclan.com/framework/router"
	"kodeclan.com/framework/utils"
)

func DefaultMethod(method string) {
	module := utils.GetCallingPackageName()
	router.DefaultMethodForModule(module, method)
}

func RegisterRequestHandler(handler http.HandlerFunc) {
	module := utils.GetCallingPackageName()
	err := router.AddModuleHandler(module, handler)

	if err != nil {
		panic(err.Error())
	}
}

func ForwardPermalinkToMethod(url, method string) {
	module := utils.GetCallingPackageName()
	router.AddStaticRoute(url, module, method)
}

func SetAsDefault() {
	module := utils.GetCallingPackageName()
	router.SetDefaultModule(module)
}
