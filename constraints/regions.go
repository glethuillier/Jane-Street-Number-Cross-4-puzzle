package constraints

// Regions-Related Constraints

import (
	"jane/common"
	"jane/row"
)

// DigitsAreOrthogonallyDifferent ensures that orthogonal digits are different
// “orthogonally adjacent cells in different regions must have different digits“
func DigitsAreOrthogonallyDifferent(
	orthogonalRegions map[rune]map[rune]struct{},
	alreadyAssigned map[rune]uint,
	currentRowOwnSymbols []rune,
	currentRowOwnDigits []uint,
) bool {
	areOrthogonallyDifferent := func(
		orthogonalRegions map[rune]map[rune]struct{},
		regionA, regionB rune,
		digitA, digitB uint,
	) bool {
		if _, ok := orthogonalRegions[regionA][regionB]; ok {
			if digitA == digitB {
				return false
			}
		}

		return true
	}

	// check digits propagated from the top row
	for k, v := range alreadyAssigned {
		for i := 0; i < len(currentRowOwnSymbols); i++ {
			if !areOrthogonallyDifferent(
				orthogonalRegions,
				currentRowOwnSymbols[i], k,
				currentRowOwnDigits[i], v,
			) {
				return false
			}
		}
	}

	// check row's own digits
	for i := 0; i < len(currentRowOwnSymbols); i++ {
		for j := i + 1; j < len(currentRowOwnSymbols); j++ {
			if !areOrthogonallyDifferent(
				orthogonalRegions,
				currentRowOwnSymbols[i], currentRowOwnSymbols[j],
				currentRowOwnDigits[i], currentRowOwnDigits[j],
			) {
				return false
			}
		}
	}

	return true
}

// TopDigitsConnectedByCurrentRowAreIdentical ensures that cells connected to the
// same regions _by the current row_ are assigned the same digits (i.e., regions
// are retrospectively connected)
//
// Example:
//
// original regions with mask:
// (top row) D*D
// (cur row) DDD
//
// updated regions:
// (top row) f*d
// (cur row) ddd
//
// assigned digits:
// (top row) 1*2
// (cur row) 222
//
// This example would be rejected as an invalid digits assignment as pseudo-region `f`
// should be assigned the same digit as region `d` because the current row connects
// these two parts, `d` and `f` belong, in fact, to the same region
func TopDigitsConnectedByCurrentRowAreIdentical(
	topRow, currentRow *row.Row,
	regions map[int][]common.Cell,
) bool {
	if len(regions) > 0 {
		for _, coordinates := range regions {
			for _, c := range coordinates {
				x, y := c.X, c.Y

				// just examine the current row cells
				// (top row: y-position 0
				//  cur row: y-position 1)
				if y == 1 {
					// if the top row and the current row, at position x, originally belong to the same
					// region and continue to belong to the same region as the current row connects them,
					// then they must be assigned to the same digit
					if topRow.OriginalRegionsWithMask[x] == currentRow.OriginalRegionsWithMask[x] &&
						topRow.Digits[x] != currentRow.Digits[x] {
						return false
					}
				}
			}
		}
	}

	return true
}
