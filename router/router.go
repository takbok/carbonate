package router

import (
	"errors"
	"net/http"
	"strings"

	"kodeclan.com/framework/utils"
)

var staticRoutes = make(map[string]string)
var defaultRoutes = make(map[string]string)
var handlers = make(map[string]http.HandlerFunc)
var defaultModule string

func AddStaticRoute(route, module, method string) {
	route = sanitizeURL(route)
	method = sanitizeURL(method)
	module = sanitizeURL(module)

	if _, exists := staticRoutes[route]; exists {
		panic(`The route '` + route + `' already exists and is mapped to '` + staticRoutes[route] + `'`)
	}

	staticRoutes[route] = module + "/" + method
	http.HandleFunc("/"+route+"/", getModuleHandler(module))
}

func SetDefaultModule(module string) {
	if defaultModule != "" {
		panic(`'` + defaultModule + `' is alrady set as the default module. Can't register '` + module + `' as the new default`)
	}

	http.HandleFunc("/", getModuleHandler(module))
	defaultModule = module
}

func AddModuleHandler(module string, handler http.HandlerFunc) error {
	if _, exists := handlers[module]; exists {
		return errors.New(`The module '` + module + `' already has a handler.`)
	}

	handlers[module] = handler
	http.HandleFunc("/"+module+"/", handler)
	return nil
}

func DefaultMethodForModule(module, method string) {
	defaultRoutes[module] = method
}

func GetParamsFromURL(url string) (string, string, []string) {
	url = getDecodedURL(url)

	components := strings.Split(url, "/")
	pkg := components[0]

	if len(components) < 2 {
		components = append(components, defaultRoutes[pkg])
	}

	module := components[0]
	method := utils.Camelize(components[1])
	args := components[2:]

	return module, method, args
}

func getStaticRoute(url string) string {
	url = sanitizeURL(url)
	return staticRoutes[url]
}

func getModuleHandler(module string) http.HandlerFunc {
	if handler, exists := handlers[module]; exists {
		return handler
	}

	panic(`No handler has been registered for the package ` + module)
}

func getDecodedURL(url string) string {
	url = sanitizeURL(url)

	if url == "" {
		url = defaultModule + "/" + defaultRoutes[defaultModule]
	} else if static, exists := staticRoutes[url]; exists {
		url = static
	}

	return url
}

func sanitizeURL(route string) string {
	return strings.Trim(route, "/ ")
}
