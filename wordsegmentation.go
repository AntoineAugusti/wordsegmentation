package wordsegmentation

import (
	"math"

	help "github.com/YafimK/wordsegmentation/helpers"
	m "github.com/YafimK/wordsegmentation/models"
)

//Corpus  lets access bigrams,
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

//Segmenator holds word segmantation functionality for specific corpus
type Segmenator struct {
	corpus     Corpus
	candidates m.Candidates
}

//NewSegmenator creates new segmentor instance that holds the language corpus and candidate table
func NewSegmenator(languageCorpus Corpus) *Segmenator {
	return &Segmenator{
		corpus: languageCorpus,
	}
}

//Segment Return a list of words that is the best segmentation of a given text.
func (s *Segmenator) Segment(text string) []string {
	return s.search(s.corpus.Clean(text), "<s>").Words
}

// Score a word in the context of the previous word.
func (s *Segmenator) score(current, previous string) float64 {
	if help.Length(previous) == 0 {
		unigramScore := s.corpus.Unigrams().ScoreForWord(current)
		if unigramScore > 0 {
			// Probability of the current word
			return unigramScore / s.corpus.Total()
		}
		// Penalize words not found in the unigrams according to their length
		return 10.0 / (s.corpus.Total() * math.Pow(10, float64(help.Length(current))))
	}
	// We've got a bigram
	unigramScore := s.corpus.Unigrams().ScoreForWord(previous)
	if unigramScore > 0 {
		bigramScore := s.corpus.Bigrams().ScoreForBigram(m.Bigram{previous, current, 0})
		if bigramScore > 0 {
			// Conditional probability of the word given the previous
			// word. The technical name is 'stupid backoff' and it's
			// not a probability distribution
			return bigramScore / s.corpus.Total() / s.score(previous, "<s>")
		}
	}

	return s.score(current, "")
}

// Find candidates for a given text and an optional previous chunk of letters.
func (s *Segmenator) findCandidates(text, prev string) <-chan m.Arrangement {
	ch := make(chan m.Arrangement)

	go func() {
		for p := range divide(text, 24) {
			prefixScore := math.Log10(s.score(p.Prefix, prev))
			arrangement := s.candidates.ForPossibility(p)
			if len(arrangement.Words) == 0 {
				arrangement = s.search(p.Suffix, p.Prefix)
				s.candidates.Add(m.Candidate{p, arrangement})
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

// Search for the best arrangement for a text in the context of a previous phrase.
func (s *Segmenator) search(text, prev string) (ar m.Arrangement) {
	if help.Length(text) == 0 {
		return m.Arrangement{}
	}

	max := -10000000.0

	// Find the best candidate by finding the best arrangement rating
	for a := range s.findCandidates(text, prev) {
		if a.Rating > max {
			max = a.Rating
			ar = a
		}
	}

	return
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
