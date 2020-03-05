package jrouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HandlerWithContext(w http.ResponseWriter, r *http.Request) {
	id := Read(r, "id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello world %s", id)))
}

func TestContext(t *testing.T) {
	t.Run("Should be able to write a variable and read it", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point/{id}", HandlerWithContext, "GET")
		r := httptest.NewRequest("GET", "/some-end-point/1", nil)
		Write(r, "id", 1)
		w := httptest.NewRecorder()
		err := jr.ServeHTTP(w, r)
		assert.Nil(t, err)
		assert.Equal(t, "Hello world 1", w.Body.String())
		assert.Equal(t, "1", Read(r, "id"))
		if status := w.Code; status != http.StatusOK {
			t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
		}
	})
}
