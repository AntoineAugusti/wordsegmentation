package wordsegmentation

import (
	"math"

	help "github.com/antoineaugusti/wordsegmentation/helpers"
	m "github.com/antoineaugusti/wordsegmentation/models"
)

var (
	corpus     Corpus
	candidates m.Candidates
)

// The corpus interface that lets access bigrams,
// unigrams, the total number of words from the corpus
// and a function to clean a string.
//
// This is the interface you will need to implement if
// you want to use a custom corpus.
type Corpus interface {
	Bigrams() *m.Bigrams
	Unigrams() *m.Unigrams
	Total() float64
	Clean(string) string
}

// Return a list of words that is the best segmentation of a given text.
func Segment(corp Corpus, text string) []string {
	corpus = corp
	return search(corpus.Clean(text), "<s>").Words
}

// Score a word in the context of the previous word.
func score(current, previous string) float64 {
	if help.Length(previous) == 0 {
		unigramScore := corpus.Unigrams().ScoreForWord(current)
		if unigramScore > 0 {
			// Probability of the current word
			return unigramScore / corpus.Total()
		} else {
			// Penalize words not found in the unigrams according to their length
			return 10.0 / (corpus.Total() * math.Pow(10, float64(help.Length(current))))
		}
	} else {
		// We've got a bigram
		unigramScore := corpus.Unigrams().ScoreForWord(previous)
		if unigramScore > 0 {
			bigramScore := corpus.Bigrams().ScoreForBigram(m.Bigram{previous, current, 0})
			if bigramScore > 0 {
				// Conditional probability of the word given the previous
				// word. The technical name is 'stupid backoff' and it's
				// not a probability distribution
				return bigramScore / corpus.Total() / score(previous, "<s>")
			}
		}

		return score(current, "")
	}
}

// Search for the best arrangement for a text in the context of a previous phrase.
func search(text, prev string) (ar m.Arrangement) {
	if help.Length(text) == 0 {
		return m.Arrangement{}
	}

	max := -10000000.0

	// Find the best candidate by finding the best arrangement rating
	for a := range findCandidates(text, prev) {
		if a.Rating > max {
			max = a.Rating
			ar = a
		}
	}

	return
}

// Find candidates for a given text and an optional previous chunk of letters.
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

// Create multiple (prefix, suffix) pairs from a text.
// The length of the prefix should not exceed the 'limit'.
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
