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
	//re := regexp.MustCompile(pattern)
	methodsParsed, err := buildMethods(methods)
	if err != nil {
		return err
	}

	//fmt.Printf("[Routes Map: %#v]\n", jr.Routes[pattern])
	if _, ok := jr.Routes[pattern]; ok {
		currentMethods := jr.Routes[pattern].Methods
		for _, mp := range methodsParsed {
			methodFound := false
			for _, m := range currentMethods {
				if m == mp {
					methodFound = true
				}
			}
			if !methodFound {
				methodsParsed = append(currentMethods, mp)
			}
		}
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

func buildMethods(methods string) ([]string, error) {
	var methodsParsed []string
	if !strings.Contains(methods, ",") {
		if !methodIsAllowed(methods) {
			return methodsParsed, MethodIsNotAllowedError
		}
		methodsParsed = append(methodsParsed, methods)
		return methodsParsed, nil
	}

	methodsList := strings.Split(methods, ",")
	for _, method := range methodsList {
		if !methodIsAllowed(method) {
			return methodsParsed, MethodIsNotAllowedError
		}
		methodsParsed = append(methodsParsed, method)
	}

	return methodsParsed, nil
}
