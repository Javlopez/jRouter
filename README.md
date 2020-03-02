[![Build Status](https://travis-ci.org/Javlopez/jrouter.svg?branch=master)](https://travis-ci.org/Javlopez/jrouter)
[![Go Report Card](https://goreportcard.com/badge/github.com/Javlopez/jrouter)](https://goreportcard.com/report/github.com/Javlopez/jrouter)
[![codecov](https://codecov.io/gh/Javlopez/jrouter/branch/master/graph/badge.svg)](https://codecov.io/gh/Javlopez/jrouter)
[![Maintainability](https://api.codeclimate.com/v1/badges/f889129ae5947f1523ec/maintainability)](https://codeclimate.com/github/Javlopez/jrouter/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/f889129ae5947f1523ec/test_coverage)](https://codeclimate.com/github/Javlopez/jrouter/test_coverage)

# jrouter

JRouter is another attemp to make easy the developers life, jrouter is trying to be simple, fast and easy Go http router.

I have started so please take a look, give me a feedback and open any issue that you have with this router.

Requeriments:

Install dependecy


```bash
go get -v https://github.com/Javlopez/jrouter
```

Basic usage


```go
//main.go
package main

import (
    "github.com/Javlopez/jrouter"
    "log"
	"net/http"
)

func main(){
    jr := New()
    jr.Handle("/some-end-point", HandlerSimpleSample, "GET")
    err := http.ListenAndServe(portNumber, jr)
    if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func HandlerSimpleSample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}
```


