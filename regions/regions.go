package regions

import (
	"jane/common"
	"jane/row"
	"slices"

	"github.com/golang-collections/collections/stack"
)

// list of characters not used for the original regions
var alphabet = []rune("abcdefghijklmnopqrstuvwxyz@NOPQRSTUVWXYZ")

// addToOrthogonalMap adds two regions, a and b, to the orthogonal map,
// which keeps track of the orthogonally adjacent regions
func addToOrthogonalMap(orthogonal map[rune]map[rune]struct{}, a, b rune) {
	if _, ok := orthogonal[a]; !ok {
		orthogonal[a] = make(map[rune]struct{})
	}
	orthogonal[a][b] = struct{}{}

	if _, ok := orthogonal[b]; !ok {
		orthogonal[b] = make(map[rune]struct{})
	}
	orthogonal[b][a] = struct{}{}
}

func getNeighbors(current common.Cell, topRow, currentRow *row.Row) []common.Cell {
	var neighbors []common.Cell

	rows := []*row.Row{topRow, currentRow}

	for _, coordinates := range []common.Coordinates{
		{
			X: current.Coordinates.X,
			Y: current.Coordinates.Y - 1,
		},
		{
			X: current.Coordinates.X,
			Y: current.Coordinates.Y + 1,
		},
		{
			X: current.Coordinates.X - 1,
			Y: current.Coordinates.Y,
		},
		{
			X: current.Coordinates.X + 1,
			Y: current.Coordinates.Y,
		},
	} {
		if coordinates.X < 0 || coordinates.Y < 0 {
			continue
		}

		if coordinates.X >= common.PuzzleSize || coordinates.Y >= 2 {
			continue
		}

		neighbors = append(neighbors, common.Cell{
			Coordinates:    coordinates,
			OriginalRegion: rows[coordinates.Y].OriginalRegionsWithMask[coordinates.X],
			UpdatedRegion:  rows[coordinates.Y].UpdatedRegions[rune(coordinates.X)],
		})

	}

	return neighbors
}

func identifyRegions(topRow, currentRow *row.Row) map[int][]common.Cell {
	s := stack.New()
	processedCells := []common.Cell{}
	regions := make(map[int][]common.Cell)
	currentRegion := 0

	for x := 0; x < common.PuzzleSize; x++ {
		currentOriginalRegion := topRow.OriginalRegionsWithMask[x]
		currentAssignedRegion := topRow.UpdatedRegions[x]

		if currentOriginalRegion == common.ShadedCell {
			continue
		}

		current := common.Cell{
			Coordinates:    common.Coordinates{X: x, Y: 0},
			OriginalRegion: currentOriginalRegion,
			UpdatedRegion:  currentAssignedRegion,
		}

		if slices.Contains(processedCells, current) {
			continue
		}

		regions[currentRegion] = []common.Cell{}

		s.Push(current)

		for {
			active := s.Pop()
			if active == nil {
				currentRegion++
				break
			}

			activeCell := active.(common.Cell)

			if activeCell.OriginalRegion == common.ShadedCell {
				continue
			}

			neighbors := getNeighbors(activeCell, topRow, currentRow)
			if len(neighbors) == 0 {
				continue
			}

			currentOriginalRegion = activeCell.OriginalRegion

			for _, neighbor := range neighbors {
				if neighbor.OriginalRegion == common.ShadedCell {
					continue
				}

				if slices.Contains(processedCells, neighbor) {
					continue
				}

				if neighbor.OriginalRegion != activeCell.OriginalRegion {
					continue
				}

				if neighbor.OriginalRegion == currentOriginalRegion {
					s.Push(neighbor)
					regions[currentRegion] = append(regions[currentRegion], neighbor)

					processedCells = append(processedCells, neighbor)
				}
			}
		}
	}

	return regions
}

// getFixedIndicesBasedOnTopRow returns the list of current row indices that correspond
// to digits propagated from the top row (i.e., not corresponding to the current row's
// own digits)
func getFixedIndicesBasedOnTopRow(regions map[int][]common.Cell) ([]int, map[int]rune) {
	unchangedRegionsIndices := []int{}
	assignedSymbolsPerRegion := make(map[int]rune)

	for regionID, v := range regions {
		for _, cell := range v {
			if cell.Y == 0 {
				assignedSymbolsPerRegion[regionID] = cell.UpdatedRegion
			} else {
				unchangedRegionsIndices = append(unchangedRegionsIndices, cell.X)
			}
		}
	}

	return unchangedRegionsIndices, assignedSymbolsPerRegion
}

// generateAssigner creates a map that associates a region
// with a given row index
func generateAssigner(
	unchangedRegionsIndices []int,
	regions map[int][]common.Cell,
	assignedSymbolsPerRegion map[int]rune,
) map[int]rune {
	assigner := make(map[int]rune)

	for _, x := range unchangedRegionsIndices {
		for regionID, v := range regions {
			for _, cell := range v {
				if cell.Y == 1 && cell.X == x {
					assigner[x] = assignedSymbolsPerRegion[regionID]
				}
			}
		}
	}

	return assigner
}

func CreateRegions(
	topRow, currentRow *row.Row,
	symbolsAlreadyMapped map[rune]uint,
) map[int][]common.Cell {
	var (
		regions                  map[int][]common.Cell
		unchangedRegionsIndices  []int
		assignedSymbolsPerRegion map[int]rune
	)

	if topRow != nil {
		regions = identifyRegions(topRow, currentRow)
		unchangedRegionsIndices, assignedSymbolsPerRegion = getFixedIndicesBasedOnTopRow(regions)
	}

	assigner := generateAssigner(
		unchangedRegionsIndices,
		regions,
		assignedSymbolsPerRegion,
	)

	alreadyProcessed := []int{}
	alreadyAssigned := make(map[rune]struct{})

	for i := 0; i < common.PuzzleSize; i++ {
		if slices.Contains(alreadyProcessed, i) {
			continue
		}

		symbol := currentRow.OriginalRegionsWithMask[i]

		if symbol == common.ShadedCell {
			continue
		}

		if value, ok := assigner[i]; ok {
			currentRow.UpdatedRegions[i] = value
		} else {
			newSymbol := getUnusedSymbol(symbolsAlreadyMapped, alreadyAssigned)
			for j := i; j < common.PuzzleSize; j++ {
				// should be assigned by assigner: skip
				_, ok := assigner[i]
				if ok {
					break
				}

				if currentRow.OriginalRegionsWithMask[j] == symbol {
					currentRow.UpdatedRegions[j] = newSymbol
					alreadyProcessed = append(alreadyProcessed, j)
				} else {
					break
				}
			}

			alreadyAssigned[newSymbol] = struct{}{}
		}
	}

	return regions
}

// getUnusedSymbol returns a symbol that has not already been used
// to designate a region
func getUnusedSymbol(
	alreadyMapped map[rune]uint,
	alreadyAssigned map[rune]struct{},
) rune {
	var (
		newSymbol rune
		offset    rune
	)

	for {
		newSymbol = alphabet[offset]

		if _, ok := alreadyMapped[newSymbol]; ok {
			offset++
			continue
		}

		if _, ok := alreadyAssigned[newSymbol]; ok {
			offset++
			continue
		}

		return newSymbol
	}
}

// GetOrthogonalRegions identifies regions that are orthogonal horizontally
// (in the same row) and, when applicable, vertically (between the current 
// row and the top row)
func GetOrthogonalRegions(topRow, bottomRow *row.Row) map[rune]map[rune]struct{} {
	orthogonal := make(map[rune]map[rune]struct{})

	rows := [][]rune{bottomRow.UpdatedRegions}

	if topRow != nil {
		rows = append(rows, topRow.UpdatedRegions)
	}

	for y := 0; y < len(rows); y++ {
		for x := 0; x < common.PuzzleSize; x++ {
			currentSymbol := rows[y][x]

			if currentSymbol == common.ShadedCell {
				continue
			}

			// horizontally orthogonal regions
			if x < common.PuzzleSize-1 &&
				rows[y][x+1] != common.ShadedCell &&
				currentSymbol != rows[y][x+1] {
				addToOrthogonalMap(orthogonal, currentSymbol, rows[y][x+1])
			}

			// vertically orthogonal regions
			if topRow != nil &&
				y == 0 &&
				rows[1][x] != common.ShadedCell &&
				currentSymbol != rows[1][x] {
				addToOrthogonalMap(orthogonal, currentSymbol, rows[1][x])
			}
		}
	}

	return orthogonal
}
