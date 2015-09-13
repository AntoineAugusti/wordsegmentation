package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	m "github.com/antoineaugusti/word-segmentation/models"
	"github.com/antoineaugusti/word-segmentation/parsers"
	"github.com/kennygrant/sanitize"
)

const (
	TOTAL = 1024908267229.0
)

var (
	bigrams    m.Bigrams
	unigrams   m.Unigrams
	candidates m.Candidates
)

func score(current, previous string) float64 {
	if length(previous) == 0 {
		unigramScore := unigrams.ScoreForWord(current)
		if unigramScore > 0 {
			return unigramScore / TOTAL
		} else {
			return 10.0 / (TOTAL * math.Pow(10, float64(length(current))))
		}
	} else {
		// We've got a bigram
		unigramScore := unigrams.ScoreForWord(previous)
		if unigramScore > 0 {
			bigramScore := bigrams.ScoreForBigram(m.Bigram{previous, current, 0})
			if bigramScore > 0 {
				return bigramScore / TOTAL / score(previous, "")
			}
		}

		return score(current, "")
	}
}

func segment(text string) []string {
	return search(cleanString(text), "<s>").Words
}

func search(text, prev string) (ar m.Arrangement) {
	if length(text) == 0 {
		return m.Arrangement{}
	}

	max := -10000000.0

	for a := range findCandidates(text, prev) {
		if a.Rating > max {
			max = a.Rating
			ar = a
		}
	}

	return
}

func findCandidates(text, prev string) <-chan m.Arrangement {
	ch := make(chan m.Arrangement)

	go func() {
		for _, p := range divide(text, 24) {
			prefixScore := math.Log10(score(p.Prefix, prev))
			arrangement := candidates.ForPossibility(p)
			if len(arrangement.Words) == 0 {
				arrangement = search(p.Suffix, p.Prefix)
				candidates.Add(m.Candidate{p, arrangement})
			}

			var slice []string
			slice = append(slice, p.Prefix)
			slice = append(slice, arrangement.Words...)
			ch <- m.Arrangement{slice, prefixScore + arrangement.Rating}
		}
		close(ch)
	}()

	return ch
}

func divide(text string, limit int) (possibilities m.Possibilities) {
	for i := 1; i <= min(length(text), limit); i++ {
		possibilities = append(possibilities, m.Possibility{text[:i], text[i:]})
	}
	return
}

func length(s string) int {
	return len([]rune(s))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	bigrams = parsers.Bigrams("data/bigrams.tsv")
	unigrams = parsers.Unigrams("data/unigrams.tsv")

	for _, l := range segment("whatistheweatherliketoday") {
		fmt.Println(l)
	}
}

func cleanString(s string) string {
	s = strings.Trim(strings.ToLower(s), " ")
	s = sanitize.Accents(s)

	// Replace certain joining characters with a dash
	s = regexp.MustCompile(`[ &_=+:]`).ReplaceAllString(s, "-")

	// Remove all other unrecognised characters
	s = regexp.MustCompile(`[^[:alnum:]-]`).ReplaceAllString(s, "")

	// Remove any multiple dashes caused by replacements above
	s = regexp.MustCompile(`[\-]+`).ReplaceAllString(s, "-")

	return s
}
