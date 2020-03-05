package jrouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HandlerSimpleSample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

func TestGeneralRouter(t *testing.T) {
	t.Run("Should be able to run new Router instance using New function", func(t *testing.T) {
		jr := New()
		assert.IsType(t, &JRouter{}, jr)
	})

	t.Run("Should be able to attach handlers", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point", HandlerSimpleSample, "POST")
		assert.IsType(t, &JRouter{}, jr)
		assert.Equal(t, 1, len(jr.Routes))
	})
	t.Run("Should reject the boostrap if the method is not allowd", func(t *testing.T) {
		jr := New()
		err := jr.Handle("/some-end-point", HandlerSimpleSample, "InvalidMethod")
		assert.Equal(t, 0, len(jr.Routes))
		assert.EqualError(t, err, "The method is not allowed")
	})

	t.Run("Should be able to group endpoint with different methods", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point", HandlerSimpleSample, "POST")
		jr.Handle("/some-end-point", HandlerSimpleSample, "GET")
		jr.Handle("/some-end-point", HandlerSimpleSample, "PUT")

		assert.IsType(t, &JRouter{}, jr)
		assert.Equal(t, 1, len(jr.Routes))
	})

	t.Run("Should be able to support different methods", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point", HandlerSimpleSample, "POST,GET")
		jr.Handle("/end-point", HandlerSimpleSample, "POST,GET, DELETE")
		jr.Handle("/end-point", HandlerSimpleSample, "POST, PUT, DELETE")

		assert.IsType(t, &JRouter{}, jr)
		assert.Equal(t, 2, len(jr.Routes))
		assert.Equal(t, 2, len(jr.Routes["/some-end-point"].Methods))
		assert.Equal(t, 4, len(jr.Routes["/end-point"].Methods))
	})

	t.Run("Should be able to run Server and dispatch the handlers", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point", HandlerSimpleSample, "POST,GET")

		r := httptest.NewRequest("GET", "/some-end-point", nil)
		w := httptest.NewRecorder()
		jr.ServeHTTP(w, r)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusNotFound, status)
		}
	})

	t.Run("Should return error if the method is not allowed", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point", HandlerSimpleSample, "POST")
		r := httptest.NewRequest("GET", "/some-end-point", nil)
		w := httptest.NewRecorder()
		err := jr.ServeHTTP(w, r)

		assert.EqualError(t, err, "http.StatusMethodNotAllowed")
		if status := w.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusMethodNotAllowed, status)
		}
	})

	t.Run("Should return status ok if the method is valid and the endpoint exists", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point", HandlerSimpleSample, "GET")
		r := httptest.NewRequest("GET", "/some-end-point", nil)
		w := httptest.NewRecorder()
		err := jr.ServeHTTP(w, r)

		assert.Nil(t, err)
		if status := w.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
		}
	})

	t.Run("Should be able to parse identifiers", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point/{id}/{name}", HandlerWithIdentifierSample, "GET")
		r := httptest.NewRequest("GET", "/some-end-point/1523/javier", nil)
		w := httptest.NewRecorder()
		err := jr.ServeHTTP(w, r)

		assert.Nil(t, err)
		assert.Equal(t, "Hello world {id:1523, name:javier}", w.Body.String())
		if status := w.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
		}
	})
}

func HandlerWithIdentifierSample(w http.ResponseWriter, r *http.Request) {
	id := Read(r, "id")
	name := Read(r, "name")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello world {id:%s, name:%s}", id, name)))
}
