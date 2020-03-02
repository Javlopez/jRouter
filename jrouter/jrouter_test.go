package jrouter

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HandlerSimpleSample(http.ResponseWriter, *http.Request) {}

func TestGeneralRouter(t *testing.T) {

	t.Run("Should be able to run new Router instance using New function", func(t *testing.T) {
		jr := New()
		assert.IsType(t, &JRouter{}, jr)
	})

	t.Run("Should be able to attach handlers", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point", HandlerSimpleSample, "POST")
		assert.IsType(t, &JRouter{}, jr)
		assert.Equal(t, len(jr.Routes), 1)
	})
	t.Run("Should reject the boostrap if the method is not allowd", func(t *testing.T) {
		jr := New()
		err := jr.Handle("/some-end-point", HandlerSimpleSample, "InvalidMethod")
		assert.Equal(t, len(jr.Routes), 0)
		assert.EqualError(t, err, "The method is not allowed")
	})

	t.Run("Should be able to group endpoint with different methods", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point", HandlerSimpleSample, "POST")
		jr.Handle("/some-end-point", HandlerSimpleSample, "GET")
		jr.Handle("/some-end-point", HandlerSimpleSample, "PUT")

		assert.IsType(t, &JRouter{}, jr)
		assert.Equal(t, len(jr.Routes), 1)
	})

	t.Run("Should be able to support different methods", func(t *testing.T) {
		jr := New()
		jr.Handle("/some-end-point", HandlerSimpleSample, "POST,GET")

		assert.IsType(t, &JRouter{}, jr)
		assert.Equal(t, 1, len(jr.Routes))
		assert.Equal(t, len(jr.Routes["/some-end-point"].Methods), 2)
	})
}
