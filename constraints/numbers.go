package constraints

// Number-Related Constraints

import (
	"jane/common"
	"jane/row"
)

// NumbersStartWithZero ensures that row numbers start with a nonzero digit
// “Numbers . . . may not begin with a 0“
func NumbersStartWithNonZero(
	currentRow *row.Row,
	alreadyAssigned map[rune]uint,
	currentRowOwnSymbols []rune,
	currentRowOwnDigits []uint,
) bool {
	numberStartsWithNonZero := func(
		symbol rune,
		alreadyAssigned map[rune]uint,
		currentRowOwnSymbols []rune,
		currentRowOwnDigits []uint,
	) bool {
		// check digits propagated from the top row
		digit, ok := alreadyAssigned[symbol]
		if ok {
			if digit == 0 {
				return false
			}
		}

		// check row's own digits
		for i, s := range currentRowOwnSymbols {
			if s == symbol {
				return currentRowOwnDigits[i] != 0
			}
		}

		return true
	}

	// first row digit
	// (only if cell has not been shaded)
	if currentRow.OriginalRegionsWithMask[0] != common.ShadedCell {
		if !numberStartsWithNonZero(
			currentRow.UpdatedRegions[0],
			alreadyAssigned,
			currentRowOwnSymbols,
			currentRowOwnDigits,
		) {
			return false
		}
	}

	// remaining digits
	for i := 0; i < common.PuzzleSize-1; i++ {
		if currentRow.OriginalRegionsWithMask[i] == common.ShadedCell {
			// assess digit next to a shaded cell 
			// (corresponding to the first digit of a number)
			if !numberStartsWithNonZero(
				currentRow.UpdatedRegions[i+1],
				alreadyAssigned,
				currentRowOwnSymbols,
				currentRowOwnDigits,
			) {
				return false
			}
		}
	}

	return true
}
