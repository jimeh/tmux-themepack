package test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertHasPrefix(t *testing.T, s, prefix string) {
	assert.Truef(t, strings.HasPrefix(s, prefix),
		"Expected \"%s\" to begin with \"%s\"", s, prefix)
}

func assertHasSuffix(t *testing.T, s, prefix string) {
	assert.Truef(t, strings.HasSuffix(s, prefix),
		"Expected \"%s\" to end with \"%s\"", s, prefix)
}
