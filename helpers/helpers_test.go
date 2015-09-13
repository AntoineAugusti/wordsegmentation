package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	assert.Equal(t, Length("abc"), 3)
	assert.Equal(t, Length(""), 0)
	assert.Equal(t, Length("helloworld"), 10)
}

func TestMin(t *testing.T) {
	assert.Equal(t, Min(-1, 0), -1)
	assert.Equal(t, Min(2, 2), 2)
	assert.Equal(t, Min(2, 3), 2)
	assert.Equal(t, Min(-1, 1), -1)
}

func TestCleanString(t *testing.T) {
	assert.Equal(t, CleanString("  abc  "), "abc")
	assert.Equal(t, CleanString("HElLo"), "hello")
	assert.Equal(t, CleanString("éàï"), "eai")
	assert.Equal(t, CleanString("a@3£4&^>?d"), "a34d")
}
