package models

type Unigram struct {
	Word   string
	Rating float64
}

type Unigrams struct {
	Data []Unigram
}

func (u *Unigram) Equals(s string) bool {
	return u.Word == s
}

func (u *Unigram) EqualsUnigram(other Unigram) bool {
	return u.Equals(other.Word)
}

func (u *Unigrams) Add(other Unigram) {
	u.Data = append(u.Data, other)
}

func (u *Unigrams) ScoreForWord(word string) float64 {
	uni := Unigram{word, 0}
	for _, unigram := range u.Data {
		if unigram.EqualsUnigram(uni) {
			return unigram.Rating
		}
	}
	return 0
}
