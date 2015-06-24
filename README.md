#Carbonate - HMVC Framework in Go

##About

Carbonate was developed for use with Google App Engine and, currently, works with the Google App Engine SDK for Go only.

Carbonate heavily follows convention over configuration. 

The URLs take the form of `/module/method/arg1/arg2/arg3`. As an example, a request to `/cart/order/58` would be handled by the `order` method of the `cart` controller and `58` will be passed as the argument.

Currently, the support for custom routes is minimal, but will be added in the future.

##Install

Carbonate can be installed via `go get`

`go get install bitbucket.org/carbonate/carbonate`


##Usage

An application is *usually* divided into multiple modules. The general convention is to have 4 files for a module:

* `handler.go` - This file is critical to setting up the module to handle requests. `handler.go` contains 2 functions only
	* `init()` - This method is responsible for registering a modules, as well as registering static routes.
	
		This function also specifies the module that should handle the default request.
	
	* `serve(http.ResponseWriter, *http.Request)` - This is the function that is called for every new incoming request.
	
		It creates a new Controller variable, assigns the proper model and views, and sends it off to a request dispatcher, which invokes the necessary controller method.
	
* `model.go` - Contains the model definition and model methods. The new model struct should embed the `BaseModel` struct.
* `view.go` - Contains the view definition and view methods. The new View struct should embed the `BaseView` struct.
* `controller.go` - Contains the controller definition and various public facing methods. The new Controller struct should embed the `BaseController` struct.

	Public facing methods should begin with an upper-case letter and must return a string only. Methods that don't follow this convention will not be publicly accessible.

Those familiar with Go will know that file are for the sake of the developers clarity only. However, the `init()` method is absolutely required to register the module and should adhere to the description above (or the example below)

##Example

Here is a simple hello world example. The module is called `hello`, and the following names should be used:

* Handler --- Contains the module initialization code
```
package hello

import (
	"net/http"

	"bitbucket.org/carbonate/carbonate"
	"bitbucket.org/carbonate/carbonate/module"
)

func init() {
	module.DefaultMethod("index")
	module.RegisterRequestHandler(serve)
	module.SetAsDefault()
}

func serve(response http.ResponseWriter, request *http.Request) {
	var controller HelloController
	controller.SetModelAndView(&controller.model, &controller.view)

	carbonate.DispatchRequestViaURL(request.URL.Path, &controller, &response, &request)
}
```

* Model --- **`HelloModel`** - Embeds `carbonate.BaseModel`
```
package hello

import "bitbucket.org/carbonate/carbonate"

type HelloModel struct {
	carbonate.BaseModel
}
```

* View --- **`HelloView`** - Embeds `carbonate.BaseView`
```
package hello

import "bitbucket.org/carbonate/carbonate"

type HelloView struct {
	carbonate.BaseView
}
```

* Controller --- **`HelloController`** - Embeds `carbonate.BaseController`
```
package hello

import "bitbucket.org/carbonate/carbonate"

type HelloController struct {
	carbonate.BaseController

	model HelloModel
	view  HelloView
}

func (c *HelloController) Index() string {
	return "Hello World"
}
```

The `Index` method above can be accessed via:

* `/hello/index` - This URL follows the `/module/method/` convention.
* `/hello` - `index` is specified as the default method for the `hello` module, hence, it is invoked when the `hello` module is invoked without a method.
* `/` - The `hello` module is specified as the default module by the `module.SetAsDefault()` function in the handler. Hence, if no module is specified, the `hello` module is invoked as the default.

More usage examples can be found at [examples repo](https://bitbucket.org/carbonate/examples)