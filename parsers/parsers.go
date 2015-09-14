package parsers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"

	m "github.com/antoineaugusti/wordsegmentation/models"
)

// Parse unigrams from a given TSV file.
func Unigrams(path string) m.Unigrams {
	jobs := make(chan string, 5000)
	results := make(chan m.Unigram, 5000)

	wg := new(sync.WaitGroup)
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go parseUnigram(jobs, results, wg)
	}

	go func() {
		readFile(path, jobs)
	}()

	// Now collect all the results
	go func() {
		wg.Wait()
		// Make sure we close the result channel when everything was processed
		close(results)
	}()

	// Add up the unigrams
	unigrams := m.NewUnigrams()
	for b := range results {
		unigrams.Add(b)
	}

	return unigrams
}

// Parse bigrams from a given TSV file.
func Bigrams(path string) m.Bigrams {
	jobs := make(chan string, 5000)
	results := make(chan m.Bigram, 5000)

	wg := new(sync.WaitGroup)
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go parseBigram(jobs, results, wg)
	}

	go func() {
		readFile(path, jobs)
	}()

	// Now collect all the results
	go func() {
		wg.Wait()
		// Make sure we close the result channel when everything was processed
		close(results)
	}()

	// Add up the bigrams
	bigrams := m.NewBigrams()
	for b := range results {
		bigrams.Add(b)
	}

	return bigrams
}

// Parse a unigram.
func parseUnigram(jobs <-chan string, results chan<- m.Unigram, wg *sync.WaitGroup) {
	defer wg.Done()
	var fields []string

	for line := range jobs {
		fields = nil
		fields = strings.Split(line, "\t")
		rating, _ := strconv.ParseFloat(fields[1], 64)

		results <- m.Unigram{fields[0], rating}
	}
}

// Parse a bigram.
func parseBigram(jobs <-chan string, results chan<- m.Bigram, wg *sync.WaitGroup) {
	defer wg.Done()
	var fields []string

	for line := range jobs {
		fields = nil
		fields = strings.Split(line, "\t")
		rating, _ := strconv.ParseFloat(fields[2], 64)

		results <- m.Bigram{fields[0], fields[1], rating}
	}
}

// Read a file and put the content in a channel.
func readFile(path string, jobs chan<- string) chan<- string {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jobs <- scanner.Text()
	}
	close(jobs)

	return jobs
}
