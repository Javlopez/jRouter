package jrouter

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func HandlerSimpleSample(http.ResponseWriter, *http.Request) {

}

func TestNewInstanceOfRouter(t *testing.T) {
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
}
