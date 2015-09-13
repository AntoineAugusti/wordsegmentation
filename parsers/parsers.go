package parsers

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	m "github.com/antoineaugusti/word-segmentation/models"
)

func Unigrams(path string) m.Unigrams {
	var fields []string
	unigrams := m.NewUnigrams()

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields = strings.Split(scanner.Text(), "\t")
		rating, _ := strconv.ParseFloat(fields[1], 64)
		unigrams.Add(m.Unigram{fields[0], rating})
	}

	return unigrams
}

func Bigrams(path string) m.Bigrams {
	var fields []string
	var words []string
	bigrams := m.NewBigrams()

	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields = strings.Split(scanner.Text(), "\t")
		words = strings.Split(fields[0], " ")
		rating, _ := strconv.ParseFloat(fields[1], 64)
		bigrams.Add(m.Bigram{words[0], words[1], rating})
	}

	return bigrams
}
