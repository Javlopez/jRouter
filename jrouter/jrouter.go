package jrouter

import (
	"errors"
	"net/http"
	"strings"
)

var AllowedMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}

var MethodIsNotAllowedError = errors.New("The method is not allowed")

//Handler is a type able to manage custom handler keeping http.ResponseWriter, *http.Request
type Handler func(http.ResponseWriter, *http.Request)

//JRouter struct
type JRouter struct {
	Routes map[string]Route
}

//Route struct
type Route struct {
	Pattern string //*regexp.Regexp
	Handler Handler
	Methods []string
}

//New function
func New() *JRouter {
	return &JRouter{
		Routes: make(map[string]Route),
	}
}

//Handle method is to create routes from handlers
func (jr *JRouter) Handle(pattern string, handler Handler, methods string) error {
	var currentMethods []string
	//re := regexp.MustCompile(pattern)
	if _, ok := jr.Routes[pattern]; ok {
		currentMethods = jr.Routes[pattern].Methods
	}

	methodsParsed, err := buildMethods(currentMethods, methods)
	if err != nil {
		return err
	}

	route := Route{Pattern: pattern, Handler: handler, Methods: methodsParsed}
	jr.Routes[pattern] = route
	return nil
}

//MethodIsAllowed validates is the method used by the request is allowed or not
func methodIsAllowed(method string) bool {
	for _, methodAllowed := range AllowedMethods {
		if methodAllowed == method {
			return true
		}
	}
	return false
}

func compileMethods(currentMethods []string, method string) []string {
	for _, m := range currentMethods {
		if m == method {
			return currentMethods
		}
	}
	return append(currentMethods, method)
}

//GET,
func buildMethods(currentMethods []string, methods string) ([]string, error) {
	var methodsParsed []string
	methodsList := strings.Split(methods, ",")

	if !strings.Contains(methods, ",") {
		methodsList = []string{methods}
	}

	for _, method := range methodsList {
		if !methodIsAllowed(method) {
			return methodsParsed, MethodIsNotAllowedError
		}
		methodsParsed = compileMethods(methodsParsed, method)
	}

	return methodsParsed, nil
}
