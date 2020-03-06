package jrouter

import (
	"errors"
	"net/http"
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
	Pattern string
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

	mb := NewMethodBuilder(currentMethods, nil)
	mb, err := mb.Add(methods)
	if err != nil {
		return err
	}

	route := Route{Pattern: pattern, Handler: handler, Methods: mb.Methods()}
	jr.Routes[pattern] = route
	return nil
}

//Get method handle GET method
func (jr *JRouter) Get(pattern string, handler Handler) error {
	return jr.Handle(pattern, handler, http.MethodGet)
}

//Post method handle POST method
func (jr *JRouter) Post(pattern string, handler Handler) error {
	return jr.Handle(pattern, handler, http.MethodPost)
}

//Put method handle PUT method
func (jr *JRouter) Put(pattern string, handler Handler) error {
	return jr.Handle(pattern, handler, http.MethodPut)
}

//Delete method handle DELETE method
func (jr *JRouter) Delete(pattern string, handler Handler) error {
	return jr.Handle(pattern, handler, http.MethodDelete)
}

//Patch method handle PATCH method
func (jr *JRouter) Patch(pattern string, handler Handler) error {
	return jr.Handle(pattern, handler, http.MethodPatch)
}

func (jr *JRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jr.context = &Context{Request: r, ResponseWriter: w}
	err := jr.dispatcher(r)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		//return err
	}

	//return nil

}

func (jr *JRouter) dispatcher(r *http.Request) error {
	for _, rt := range jr.Routes {
		uParser := NewURLParser()
		uParser.Analyze(rt.Pattern)

		if matches := uParser.PatternMatcher.FindStringSubmatch(jr.context.URL.Path); len(matches) > 0 {

			parameters := matches[1:]
			if parameters != nil {
				for i, p := range uParser.Params {
					Write(r, p.Param, parameters[i])
				}
			}

			mb := NewMethodBuilder(nil, rt.Methods)
			if !mb.MethodIsAllowed(jr.context.Method) {
				return errors.New("http.StatusMethodNotAllowed")
			}
			rt.Handler(jr.context.ResponseWriter, jr.context.Request)
		}
	}

	return nil
}
