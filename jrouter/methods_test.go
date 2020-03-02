package jrouter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethods(t *testing.T) {
	t.Run("Should be able to set and receive slice of methods", func(t *testing.T) {
		methods := []string{"GET"}
		mb := NewMethodBuilder(methods)
		assert.Equal(t, methods, mb.Methods())
	})

	t.Run("Should support add new methods", func(t *testing.T) {
		methods := []string{}
		mb := NewMethodBuilder(methods)
		mb.Add("GET")
		assert.Equal(t, []string{"GET"}, mb.Methods())
	})

	t.Run("Should support merge methods avoiding duplicated", func(t *testing.T) {
		methods := []string{"GET", "POST"}
		mb := NewMethodBuilder(methods)
		mb.Add("GET")
		assert.Equal(t, len([]string{"POST", "GET"}), len(mb.Methods()))
	})
}
