package main

import (
	"fmt"
	"math"

	help "github.com/antoineaugusti/word-segmentation/helpers"
	m "github.com/antoineaugusti/word-segmentation/models"
	"github.com/antoineaugusti/word-segmentation/parsers"
)

const (
	TOTAL = 1024908267229.0
)

var (
	bigrams    m.Bigrams
	unigrams   m.Unigrams
	candidates m.Candidates
)

// Load unigrams and bigrams on initialization
func init() {
	done := make(chan int)

	go func() {
		bigrams = parsers.Bigrams("data/bigrams.tsv")
		done <- 1
	}()
	go func() {
		unigrams = parsers.Unigrams("data/unigrams.tsv")
		done <- 1
	}()

	<-done
	<-done
}

func main() {
	for _, l := range Segment("whatistheweatherliketoday") {
		fmt.Println(l)
	}
}

func Segment(text string) []string {
	return search(help.CleanString(text), "<s>").Words
}

func score(current, previous string) float64 {
	if help.Length(previous) == 0 {
		unigramScore := unigrams.ScoreForWord(current)
		if unigramScore > 0 {
			return unigramScore / TOTAL
		} else {
			return 10.0 / (TOTAL * math.Pow(10, float64(help.Length(current))))
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

func search(text, prev string) (ar m.Arrangement) {
	if help.Length(text) == 0 {
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
		for p := range divide(text, 24) {
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

func divide(text string, limit int) <-chan m.Possibility {
	ch := make(chan m.Possibility)
	bound := help.Min(help.Length(text), limit)

	go func() {
		for i := 1; i <= bound; i++ {
			ch <- m.Possibility{Prefix: text[:i], Suffix: text[i:]}
		}
		close(ch)
	}()

	return ch
}
