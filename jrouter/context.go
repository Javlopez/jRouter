package jrouter

import (
	"net/http"
)

//RequestContext is storage data to share between router and handler
var (
	RequestContext = make(map[*http.Request]map[string]interface{})
)

//Write is designed to set data into RequestContext variable
func Write(r *http.Request, name string, value interface{}) {
	if RequestContext[r] == nil {
		RequestContext[r] = make(map[string]interface{})
	}
	RequestContext[r][name] = value
}

//Read is designed get data from RequestContext variable
func Read(r *http.Request, name string) interface{} {
	return RequestContext[r][name]
}
