package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForPossibility(t *testing.T) {
	candidates := Candidates{}
	arrangement := Arrangement{[]string{"a", "b"}, 42}
	possibility := Possibility{"prefix", "suffix"}
	candidate := Candidate{possibility, arrangement}
	candidates.Add(candidate)

	assert.Equal(t, candidates.ForPossibility(possibility), arrangement)
	assert.Equal(t, len(candidates.ForPossibility(Possibility{"not", "found"}).Words), 0)
}
