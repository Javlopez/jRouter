package jrouter

import (
	"errors"
	"net/http"
	"strings"
)

var MethodIsNotAllowedError = errors.New("The method is not allowed")

var AllowedMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}

//MethodBuilder struct
type MethodBuilder struct {
	methods []string
}

//NewMethodBuilder func
func NewMethodBuilder(methods []string) *MethodBuilder {
	return &MethodBuilder{methods: methods}
}

//Add method
func (mb *MethodBuilder) Add(methods string) (*MethodBuilder, error) {
	//var methodsParsed []string
	methodsList := strings.Split(methods, ",")

	if !strings.Contains(methods, ",") {
		methodsList = []string{methods}
	}

	for _, method := range methodsList {

		method = strings.TrimSpace(method)
		if !mb.methodIsAllowed(method) {
			return mb, MethodIsNotAllowedError
		}
		mb.methods = mb.compileMethods(mb.methods, method)
	}
	return mb, nil
}

func (mb *MethodBuilder) methodIsAllowed(method string) bool {
	for _, methodAllowed := range AllowedMethods {
		if methodAllowed == method {
			return true
		}
	}
	return false
}

func (mb *MethodBuilder) compileMethods(currentMethods []string, method string) []string {
	for _, m := range currentMethods {
		if m == method {
			return currentMethods
		}
	}
	return append(currentMethods, method)
}

func (mb *MethodBuilder) Methods() []string {
	return mb.methods
}
