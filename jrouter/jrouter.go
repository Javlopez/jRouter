package jrouter

import (
	"net/http"
)

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

	mb := NewMethodBuilder(currentMethods)
	mb, err := mb.Add(methods)
	if err != nil {
		return err
	}

	route := Route{Pattern: pattern, Handler: handler, Methods: mb.Methods()}
	jr.Routes[pattern] = route
	return nil
}
