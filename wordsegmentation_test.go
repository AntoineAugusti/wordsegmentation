package wordsegmentation

import (
	"testing"

	c "github.com/AntoineAugusti/wordsegmentation/corpus"
	"github.com/stretchr/testify/assert"
)

func TestSegment(t *testing.T) {

	expected := []string{"what", "is", "the", "weather", "like", "today"}
	englishCorpus := c.NewEnglishCorpus()
	segmentaor := NewSegmenator(englishCorpus)
	assert.Equal(t, segmentaor.Segment(englishCorpus, "WhatIsTheWeatherliketoday? "), expected)
}
