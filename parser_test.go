package jrouter

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {

	t.Run("Test simple endpoint", func(t *testing.T) {
		url := "/something/"
		p := NewURLParser()
		p.Analyze(url)
		re := regexp.MustCompile(url)
		assert.Equal(t, re, p.PatternMatcher)
	})

	t.Run("Test basic params", func(t *testing.T) {
		url := "/something/{id}/{name}"
		p := NewURLParser()
		p.Analyze(url)
		re := regexp.MustCompile("/something/([a-zA-Z\\d+]+)/([a-zA-Z\\d+]+)")
		assert.Equal(t, re, p.PatternMatcher)
	})
}
