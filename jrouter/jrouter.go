package jrouter

import (
	"net/http"
)

//Handler is a type able to manage custom handler keeping http.ResponseWriter, *http.Request
type Handler func(http.ResponseWriter, *http.Request)

//JRouter struct
type JRouter struct {
	Routes []Route
}

//Route struct
type Route struct {
	Pattern string //*regexp.Regexp
	Handler Handler
	Methods string
}

//New function
func New() *JRouter {
	return &JRouter{}
}

//Handle method is to create routes from handlers
func (jr *JRouter) Handle(pattern string, handler Handler, methods string) {
	//re := regexp.MustCompile(pattern)
	route := Route{Pattern: pattern, Handler: handler, Methods: methods}
	jr.Routes = append(jr.Routes, route)
}

/*



jr := jrouter.New()
r.Handle("endpoint", handler, "POST, GET")
r.Handle("endpoint", handler, "POST, GET")
r.Serve(8080)
*/
