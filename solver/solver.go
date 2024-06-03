package solver

import (
	"jane/common"
	"jane/constraints"
	"jane/mask"
	"jane/regions"
	"jane/row"
	"sync"
)

func Solve() {
	var wg sync.WaitGroup

	solveCandidate(&wg, &Candidate{})

	wg.Wait()
}

func isValidRow(
	currentRow, topRow *row.Row,
	candidate *Candidate,
	orthogonalRegions map[rune]map[rune]struct{},
	currentRowAllRegions map[int][]common.Cell,
	currentRowOwnRegions []rune,
	currentRowOwnDigits []uint,
) bool {
	// ensure that adjacent cells in different regions have different digits
	if !constraints.DigitsAreOrthogonallyDifferent(
		orthogonalRegions,
		candidate.ExistingRegions,
		currentRowOwnRegions,
		currentRowOwnDigits,
	) {
		return false
	}

	// ensure that numbers do not start with 0
	if !constraints.NumbersStartWithNonZero(
		currentRow,
		candidate.ExistingRegions,
		currentRowOwnRegions,
		currentRowOwnDigits,
	) {
		return false
	}

	// assign digits to each region
	currentRow.ResetAssignments()
	for k, v := range candidate.ExistingRegions {
		currentRow.AssignDigit(k, v)
	}
	for i := 0; i < len(currentRowOwnRegions); i++ {
		currentRow.AssignDigit(currentRowOwnRegions[i], uint(currentRowOwnDigits[i]))
	}

	// ensure that parts of top row regions connected by the current rows
	// are assigned the same digits
	if !constraints.TopDigitsConnectedByCurrentRowAreIdentical(
		topRow, currentRow,
		currentRowAllRegions,
	) {
		return false
	}

	// ensure that the row numbers satisfy the row clue
	return constraints.NumbersSatisfyTheRowClue(candidate.RowID, currentRow.GetNumbers())
}

func solveCandidate(wg *sync.WaitGroup, candidate *Candidate) {
	defer wg.Done()

	rowID := candidate.RowID

	// guard clause: end of exploration for this branch
	// (solution found)
	if rowID == common.PuzzleSize {
		return
	}

	// if applicable, instantiate the top row
	// (i.e., the row immediately above the current row)
	var topRow *row.Row
	if rowID != 0 {
		topRow = candidate.Rows[rowID-1]
	}

	// check each configuration of shaded cells
	for maskID := 0; maskID < mask.GetMasksCount(); maskID++ {

		// ensure that the position of the shaded cells is compatible
		// with the top row ones (nonadjacent shaded cells)
		currentMask := mask.GetMask(maskID)
		if topRow != nil && !topRow.IsMaskCompatibleWith(currentMask) {
			continue
		}

		currentRow := row.GenerateRow(rowID, currentMask)

		// create regions based on the original regions and the current
		// configuration of shaded cells
		currentRowAllRegions := regions.CreateRegions(
			topRow, currentRow,
			candidate.ExistingRegions,
		)
		orthogonalRegions := regions.GetOrthogonalRegions(topRow, currentRow)
		currentRowOwnRegions := currentRow.GetRegionsToAssign(candidate.ExistingRegions)

		for currentRowOwnDigits := range generateDigits(len(currentRowOwnRegions)) {
			if isValidRow(
				currentRow, topRow,
				candidate,
				orthogonalRegions,
				currentRowAllRegions,
				currentRowOwnRegions,
				currentRowOwnDigits,
			) {
				// the current row satisfies all constraints

				// only display the first valid row for each level and
				// display all solutions (theoretically, as there is
				// in fact only one solution)
				if rowID == common.PuzzleSize-1 || isGreaterThanCounter(int64(rowID)) {
					setCounter(int64(rowID))
					displayRows(rowID, candidate, currentRow)
				}

				// propagate the candidate constraints:
				// (a) augment the candidate by adding the current row
				augmentedCandidate := AugmentCandidate(rowID, candidate, currentRow)

				// (b) try solving the augmented candidate in a new goroutine
				wg.Add(1)
				go solveCandidate(wg, augmentedCandidate)

				SaveCandidate(rowID, candidate, currentRow)
			}
		}
	}
}

// generateDigits generates the Cartesian product of
// digits from 0 to 9 to be assigned to new regions
func generateDigits(length int) <-chan []uint {
	ch := make(chan []uint)

	go func() {
		defer close(ch)
		digits := make([]uint, length)

		var generate func(int)

		generate = func(index int) {
			if index == length {
				clone := make([]uint, length)
				copy(clone, digits)
				ch <- clone
				return
			}

			for i := uint(0); i <= 9; i++ {
				digits[index] = i
				generate(index + 1)
			}
		}

		generate(0)
	}()

	return ch
}
