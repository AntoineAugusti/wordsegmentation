package wordsegmentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSegment(t *testing.T) {
	assert.Equal(t, Segment("WhatIsTheWeatherliketoday? "), []string{"what", "is", "the", "weather", "like", "today"})
}
