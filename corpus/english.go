package corpus

import (
	"os"
	"path"

	help "github.com/antoineaugusti/wordsegmentation/helpers"
	m "github.com/antoineaugusti/wordsegmentation/models"
	"github.com/antoineaugusti/wordsegmentation/parsers"
)

type EnglishCorpus struct {
	unigrams m.Unigrams
	bigrams  m.Bigrams
}

// Create a new EnglishCorpus by loading unigrams
// and bigrams from TSV files.
func NewEnglishCorpus() EnglishCorpus {
	var bigrams m.Bigrams
	var unigrams m.Unigrams

	// Load unigrams and bigrams from data files.
	done := make(chan int)
	goPath := path.Clean(os.Getenv("GOPATH"))
	dataPath := goPath + "/src/github.com/antoineaugusti/wordsegmentation/data/"

	go func(dataPath string) {
		bigrams = parsers.Bigrams(dataPath + "english/bigrams.tsv")
		done <- 1
	}(dataPath)
	go func(dataPath string) {
		unigrams = parsers.Unigrams(dataPath + "english/unigrams.tsv")
		done <- 1
	}(dataPath)

	<-done
	<-done

	return EnglishCorpus{unigrams, bigrams}
}

// Get bigrams from the corpus.
func (corpus EnglishCorpus) Bigrams() *m.Bigrams {
	return &corpus.bigrams
}

// Get unigrams from the corpus.
func (corpus EnglishCorpus) Unigrams() *m.Unigrams {
	return &corpus.unigrams
}

// Get the total number of words in the corpus.
func (corpus EnglishCorpus) Total() float64 {
	return 1024908267229.0
}

// Clean a string from special characters.
func (corpus EnglishCorpus) Clean(s string) string {
	return help.CleanString(s)
}
