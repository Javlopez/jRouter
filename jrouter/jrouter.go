package jrouter

import (
	"errors"
	"net/http"
	"regexp"
)

//Handler is a type able to manage custom handler keeping http.ResponseWriter, *http.Request
type Handler func(http.ResponseWriter, *http.Request)

//Context is a struct to manage the context
type Context struct {
	http.ResponseWriter
	*http.Request
}

//JRouter struct
type JRouter struct {
	Routes  map[string]Route
	context *Context
}

//Route struct
type Route struct {
	Pattern *regexp.Regexp
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
	re := regexp.MustCompile(pattern)
	if _, ok := jr.Routes[pattern]; ok {
		currentMethods = jr.Routes[pattern].Methods
	}

	mb := NewMethodBuilder(currentMethods, nil)
	mb, err := mb.Add(methods)
	if err != nil {
		return err
	}

	route := Route{Pattern: re, Handler: handler, Methods: mb.Methods()}
	jr.Routes[pattern] = route
	return nil
}

func (jr *JRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	jr.context = &Context{Request: r, ResponseWriter: w}
	err := jr.dispatcher()
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return err
	}

	return nil

}

func (jr *JRouter) dispatcher() error {
	for _, rt := range jr.Routes {
		if matches := rt.Pattern.FindStringSubmatch(jr.context.URL.Path); len(matches) > 0 {

			mb := NewMethodBuilder(nil, rt.Methods)
			if !mb.MethodIsAllowed(jr.context.Method) {
				return errors.New("http.StatusMethodNotAllowed")
			}
			rt.Handler(jr.context.ResponseWriter, jr.context.Request)
		}
	}

	return nil
}
