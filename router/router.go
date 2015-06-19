// The Takbok router package contains methods to route various real world URLs
// to functions and methods, including static routing
package router

import (
	"net/http"
	"strings"

	"kodeclan.com/framework/utils"
)

// A map of all the static routes
var staticRoutes = make(map[string]string)

// A map of all the default routes
var defaultRoutes = make(map[string]string)

// A map of all the handler functions
var handlers = make(map[string]http.HandlerFunc)

// Stores the name of the default module to be used when "/" is requested
var defaultModule string

// Add a static route to the routing table
// Maps a real world URL to a controller method of a module
func AddStaticRoute(route, module, method string) {
	route = sanitizeURL(route)
	method = sanitizeURL(method)
	module = sanitizeURL(module)

	if _, exists := staticRoutes[route]; exists {
		panic(`The route '` + route + `' already exists and is mapped to '` + staticRoutes[route] + `'`)
	}

	// Register the static route to the module handler
	http.HandleFunc("/"+route+"/", getModuleHandler(module))
	staticRoutes[route] = module + "/" + method
}

// Specify the default module for the application
// Requests to "/" will be redirected to this module
func SetDefaultModule(module string) {
	if defaultModule != "" {
		panic(`'` + defaultModule + `' is alrady set as the default module. Can't register '` + module + `' as the new default`)
	}

	// Register the module handler to the application root
	http.HandleFunc("/", getModuleHandler(module))
	defaultModule = module
}

// Add the handler method for the module.
// All incoming requests to a module get directed to this method.
func AddModuleHandler(module string, handler http.HandlerFunc) {
	if _, exists := handlers[module]; exists {
		panic(`The module '` + module + `' already has a handler.`)
	}

	// Register the handler to the module name
	http.HandleFunc("/"+module+"/", handler)
	handlers[module] = handler
}

// Specify the default controller method for the module
func DefaultMethodForModule(module, method string) {
	defaultRoutes[module] = method
}

// Get the name of the module, the method, and a list of arguments from a URL
func GetParamsFromURL(url string) (string, string, []string) {
	url = getDecodedURL(url)

	// Tokenize the URL
	components := strings.Split(url, "/")
	pkg := components[0]

	if len(components) < 2 {
		// Add the default method for the module as the route,
		// if the URL doesn't specify any method
		components = append(components, defaultRoutes[pkg])
	}

	module := components[0]
	method := utils.Camelize(components[1])
	args := components[2:]

	return module, method, args
}

// Get the static route mapped to a URL
func getStaticRoute(url string) string {
	url = sanitizeURL(url)
	return staticRoutes[url]
}

// Get the handler function for a module.
func getModuleHandler(module string) http.HandlerFunc {
	if handler, exists := handlers[module]; exists {
		return handler
	}

	panic(`No handler has been registered for the package ` + module)
}

// Get the proper URL relative to a entry in the routing table
func getDecodedURL(url string) string {
	url = sanitizeURL(url)

	if url == "" {
		// if the URL is empty, then return the default module and method
		url = defaultModule + "/" + defaultRoutes[defaultModule]
	} else if static, exists := staticRoutes[url]; exists {
		// otherwise, check if a static route exists
		url = static
	}

	return url
}

// Sanitize a string by removing leading and trailing slashes
// and any type of whitespaces
func sanitizeURL(route string) string {
	return strings.Trim(route, "/ ")
}
