package wordsegmentation

import (
	"testing"
	"fmt"
	c "github.com/AntoineAugusti/wordsegmentation/corpus"
	"github.com/stretchr/testify/assert"
)

func TestSegmenator_Segment(t *testing.T) {
	type args struct {
		Text string
	}
	englishCorpus := c.NewEnglishCorpus()

	tests := []struct {
		name string
		s    *Segmenator
		args args
		want []string
	}{
		{"basic", NewSegmenator(englishCorpus), args{"WhatIsTheWeatherliketoday?"}, []string{"what", "is", "the", "weather", "like", "today"}},
		{"basic with spaces", NewSegmenator(englishCorpus), args{"click me next"}, []string{"click", "me", "next"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Segment(tt.args.Text)
			assert.Equalf(t, tt.want, got, fmt.Sprintf("Segmenator.Segment() = %v, want %v\n", got, tt.want))
		})
	}
}
