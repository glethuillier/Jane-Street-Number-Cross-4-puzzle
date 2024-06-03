package solver

import "jane/row"

type Candidate struct {
	RowID           uint
	Rows            []*row.Row
	ExistingRegions map[rune]uint
}

// AugmentCandidate extends a given candidate with a new valid row and updates
// its existing regions accordingly
func AugmentCandidate(rowID uint, candidate *Candidate, currentRow *row.Row) *Candidate {
	augmentedCandidate := Candidate{
		RowID:           rowID + 1,
		Rows:            make([]*row.Row, rowID+1),
		ExistingRegions: make(map[rune]uint),
	}

	for i := range rowID {
		augmentedCandidate.Rows[i] = row.Copy(candidate.Rows[i])

		for k, v := range candidate.Rows[i].GetRegionToDigitMap() {
			augmentedCandidate.ExistingRegions[k] = v
		}
	}

	augmentedCandidate.Rows[rowID] = row.Copy(currentRow)
	for k, v := range currentRow.GetRegionToDigitMap() {
		augmentedCandidate.ExistingRegions[k] = v
	}

	return &augmentedCandidate
}
