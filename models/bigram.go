package models

import "fmt"

type Bigram struct {
	First  string
	Second string
	Rating float64
}

type Bigrams struct {
	data map[string]float64
}

func (b *Bigram) GetKey() string {
	return fmt.Sprintf("%s#%s", b.First, b.Second)
}

func NewBigrams() Bigrams {
	return Bigrams{data: make(map[string]float64)}
}

func (b *Bigrams) Add(other Bigram) {
	b.data[other.GetKey()] = other.Rating
}

func (b *Bigrams) ScoreForBigram(other Bigram) float64 {
	score, has := b.data[other.GetKey()]
	if !has {
		return 0
	}
	return score
}
