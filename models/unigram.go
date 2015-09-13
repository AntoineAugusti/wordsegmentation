package models

type Unigram struct {
	Word   string
	Rating float64
}

type Unigrams struct {
	data map[string]float64
}

// Create a new unigrams collection.
func NewUnigrams() Unigrams {
	return Unigrams{data: make(map[string]float64)}
}

// Add a unigram to the collection.
func (u *Unigrams) Add(other Unigram) {
	u.data[other.Word] = other.Rating
}

// Find the score for a given string. If the string
// was not found, the score will be 0.
func (u *Unigrams) ScoreForWord(word string) float64 {
	score, has := u.data[word]
	if !has {
		return 0
	}
	return score
}
