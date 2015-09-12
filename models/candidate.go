package models

type Candidate struct {
	P Possibility
	A Arrangement
}

type Candidates struct {
	data []Candidate
}

func (c *Candidates) Add(other Candidate) {
	c.data = append(c.data, other)
}

func (c *Candidates) ForPossibility(p Possibility) Arrangement {
	for _, candidate := range c.data {
		if candidate.P.Prefix == p.Prefix && candidate.P.Suffix == p.Suffix {
			return candidate.A
		}
	}
	return Arrangement{}
}
