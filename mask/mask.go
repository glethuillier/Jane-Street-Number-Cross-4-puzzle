package mask

import (
	"fmt"
	"jane/common"
)

// masks represent valid configurations of shaded cells
var masks map[int][]bool

// isValidConfigurationOfShadedCells ensures that a generated mask complies
// with the constraints of the puzzle:
//
// “Shading must be “sparse”: that is, no two shaded cells may share an edge.“
// “Numbers must be at least two digits long“
func isValidConfigurationOfShadedCells(slice []bool) bool {
	// first-digit cell cannot be followed by a shaded cell;
	// otherwise, a number could be formed of one-digit
	// 01---
	if !slice[0] && slice[1] {
		return false
	}

	// last-digit cell cannot be preceded by a shaded cell;
	// otherwise, a number could be formed of one-digit
	// ---10
	if !slice[len(slice)-1] && slice[len(slice)-2] {
		return false
	}

	for i := 1; i < len(slice); i++ {
		// no single-digit cell should be surrounded by two shaded cells;
		// otherwise, a number could be formed of one-digit
		// ---101---
		if i >= 2 && slice[i-2] && !slice[i-1] && slice[i] {
			return false
		}

		// shaded cells should not be adjacent;
		// otherwise, the shaded cells would not be sparse
		// ---11---
		if slice[i-1] && slice[i] {
			return false
		}
	}

	return true
}

func init() {
	masks = make(map[int][]bool)
	configurations := generateCartesianProduct(common.PuzzleSize)

	maskID := 0
	for _, p := range configurations {
		if !isValidConfigurationOfShadedCells(p) {
			continue
		}

		_, ok := masks[maskID]
		if !ok {
			masks[maskID] = []bool{}
		}

		masks[maskID] = p
		maskID++
	}

	fmt.Printf("Generated %d masks (shaded cells)\n", len(masks))
}

func generateCartesianProduct(length int) [][]bool {
	total := 1 << uint(length)
	configurations := make([][]bool, total)

	for i := 0; i < total; i++ {
		product := make([]bool, length)
		for j := 0; j < length; j++ {
			product[j] = (i>>j)&1 == 1
		}

		configurations[i] = product
	}

	return configurations
}

func GetMask(maskID int) []bool {
	return masks[maskID]
}

func GetMasksCount() int {
	return len(masks)
}
