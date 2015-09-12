package models

type Bigram struct {
	First  string
	Second string
	Rating float64
}

type Bigrams struct {
	Data []Bigram
}

func (b *Bigram) Equals(first, second string) bool {
	return b.First == first && b.Second == second
}

func (b *Bigram) EqualsBigram(other Bigram) bool {
	return b.Equals(other.First, other.Second)
}

func (b *Bigrams) Add(other Bigram) {
	b.Data = append(b.Data, other)
}

func (b *Bigrams) ScoreForBigram(other Bigram) float64 {
	for _, bigram := range b.Data {
		if bigram.EqualsBigram(other) {
			return bigram.Rating
		}
	}
	return 0
}
