package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScoreForWord(t *testing.T) {
	collection := NewUnigrams()
	unigram := Unigram{"foo", 22}
	collection.Add(unigram)

	assert.Equal(t, collection.ScoreForWord(unigram.Word), unigram.Rating)
	assert.Equal(t, collection.ScoreForWord("notFound"), 0.0)
}
