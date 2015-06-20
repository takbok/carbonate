/**
 * The Takbok framework is an HMVC framework for Go.
 *
 * The framework package contains the Model, View and Controller, along with the
 * Request dispatcher and the Template handler
 *
 * Everything is organized into multiple modules.
 *
 * It's advisable to have a "modules" folder in the same directory as the framework
 * folder. The "modules" folder should have all the modules, each under a directory
 * names after the module, and each module having a similar Go package name.
 *
 * Each module should also have a directory called "templates" with ONE OR MORE
 * files with the extension ".tmpl"
 *
 * The structure should be like
 * + framework
 *   - (this framework)
 * + modules
 *   + module1
 *     + templates
 *       - template1.tmpl
 *     - controller.go
 *     - handler.go
 *     - model.go
 *     - view.go
 *
 *   + module2
 *     + tempaltes
 *       - template1.tmpl
 *     - controller.go
 *     - handler.go
 *     - model.go
 *     - view.go
 *
 * The controller.go, view.go and model.go will contain the code for the
 * Controller, View and Model respectively.
 *
 * handler.go should ideally contain two functions only.
 * - serve (or any other name) - a function of http.HandlerFunc type
 *      This function should declare a new module native controller, then
 *      call the SetModelAndView method with the BaseModel and BaseView of the
 *      controller as the parameters.
 *
 *      Finally, it should dispatch the request via the DispatchRequestViaURL function
 *
 *      Example:
 *
 *        func serve(response http.ResponseWriter, request *http.Request) {
 *            var c NewController
 *            c.SetModelAndView(&c.model, &c.view)
 *
 *            framework.DispatchRequestViaURL(request.URL.Path, &c, &response, &request)
 *        }
 *
 * - init - This function should register the default method for the module,
 *      and the http.HandlerFunc of the module.
 *
 *      Optionally, it may also specify if the module is the default module for
 *      the application, and also register any static routes that the module should handle
 *
 *      Example:
 *
 *        func init() {
 *            module.DefaultMethod("welcome") // By default, the "Welcome" method will be called
 *            module.RegisterRequestHandler(serve)
 *            module.SetAsDefault() // If this is the default module
 *
 *            // Any incoming request to "/hello/" should be handled by the "login" method
 *            module.ForwardPermalinkToMethod("/home/", "login")
 *
 *            // Any incoming request to "/bye/" should be handled by the "logout" method
 *            module.ForwardPermalinkToMethod("/bye/", "logout")
 *        }
 *
 */
package framework
