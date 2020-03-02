package jrouter

import (
	"errors"
	"net/http"
	"strings"
)

var MethodIsNotAllowedError = errors.New("The method is not allowed")

var DefaultMethodsAllowed = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}

//MethodBuilder struct
type MethodBuilder struct {
	methods        []string
	methodsAllowed []string
}

//NewMethodBuilder func
func NewMethodBuilder(methodsBase, methodsAllowed []string) *MethodBuilder {
	if methodsAllowed == nil {
		methodsAllowed = DefaultMethodsAllowed
	}
	return &MethodBuilder{methods: methodsBase, methodsAllowed: methodsAllowed}
}

//Add method
func (mb *MethodBuilder) Add(methods string) (*MethodBuilder, error) {

	methodsList := strings.Split(methods, ",")
	if !strings.Contains(methods, ",") {
		methodsList = []string{methods}
	}

	for _, method := range methodsList {

		method = strings.TrimSpace(method)
		if !mb.MethodIsAllowed(method) {
			return mb, MethodIsNotAllowedError
		}
		mb.methods = mb.compileMethods(mb.methods, method)
	}
	return mb, nil
}

//MethodIsAllowed method
func (mb *MethodBuilder) MethodIsAllowed(method string) bool {
	for _, methodAllowed := range mb.methodsAllowed {
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

//Methods method
func (mb *MethodBuilder) Methods() []string {
	return mb.methods
}
