package models

type Candidate struct {
	P Possibility
	A Arrangement
}

type Candidates struct {
	data []Candidate
}

// Add another Candidate to the collection.
func (c *Candidates) Add(other Candidate) {
	c.data = append(c.data, other)
}

// Find an Arrangement for a Possibility. If the Possibility was
// not found, an empty Arrangement will be given.
func (c *Candidates) ForPossibility(p Possibility) Arrangement {
	for _, candidate := range c.data {
		if candidate.P.Prefix == p.Prefix && candidate.P.Suffix == p.Suffix {
			return candidate.A
		}
	}
	return Arrangement{}
}
