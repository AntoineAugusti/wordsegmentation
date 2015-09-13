package models

type Unigram struct {
	Word   string
	Rating float64
}

type Unigrams struct {
	data map[string]float64
}

func NewUnigrams() Unigrams {
	return Unigrams{data: make(map[string]float64)}
}

func (u *Unigrams) Add(other Unigram) {
	u.data[other.Word] = other.Rating
}

func (u *Unigrams) ScoreForWord(word string) float64 {
	score, has := u.data[word]
	if !has {
		return 0
	}
	return score
}
