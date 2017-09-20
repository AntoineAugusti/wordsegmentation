package corpus

import (
	"os"
	"path/filepath"

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
	for _, component := range filepath.SplitList(os.Getenv("GOPATH")) {
		component = filepath.Clean(component)
		dataPath := filepath.Join(component, "src", "github.com", "antoineaugusti", "wordsegmentation", "data")
		if _, err := os.Stat(dataPath); os.IsNotExist(err) {
			continue
		}
		go func(dataPath string) {
			bigrams = parsers.Bigrams(filepath.Join(dataPath, "english", "bigrams.tsv"))
			done <- 1
		}(dataPath)
		go func(dataPath string) {
			unigrams = parsers.Unigrams(filepath.Join(dataPath, "english", "unigrams.tsv"))
			done <- 1
		}(dataPath)
		<-done
		<-done
		break
	}

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
