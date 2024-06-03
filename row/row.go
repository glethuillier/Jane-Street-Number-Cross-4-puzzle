package row

import (
	"encoding/json"
	"fmt"
	"jane/common"
	"os"
	"path/filepath"
	"slices"
)

type Row struct {
	OriginalRegionsWithMask []rune
	UpdatedRegions          []rune
	Digits                  []uint
	regionToDigitMap        map[rune]uint
}

var rows map[uint]*Row

func Copy(original *Row) *Row {
	new := Row{}

	new.OriginalRegionsWithMask = make([]rune, len(original.OriginalRegionsWithMask))
	copy(new.OriginalRegionsWithMask, original.OriginalRegionsWithMask)

	new.UpdatedRegions = make([]rune, len(original.UpdatedRegions))
	copy(new.UpdatedRegions, original.UpdatedRegions)

	new.Digits = make([]uint, len(original.Digits))
	copy(new.Digits, original.Digits)

	new.regionToDigitMap = make(map[rune]uint)
	for k, v := range original.regionToDigitMap {
		new.regionToDigitMap[k] = v
	}

	return &new
}

type Regions []struct {
	Region      string `json:"region"`
	Coordinates []struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"coordinates"`
}

func init() {
	rows = make(map[uint]*Row, common.PuzzleSize)

	content, err := os.ReadFile(filepath.Join(
		common.GetBasePath(),
		"../puzzle.json",
	))
	if err != nil {
		panic(err)
	}

	var segments Regions
	err = json.Unmarshal(content, &segments)
	if err != nil {
		panic(err)
	}

	for _, segment := range segments {
		for _, cell := range segment.Coordinates {
			if _, ok := rows[uint(cell.Y)]; !ok {
				rows[uint(cell.Y)] = &Row{
					UpdatedRegions:          make([]rune, common.PuzzleSize),
					OriginalRegionsWithMask: make([]rune, common.PuzzleSize),
					Digits:                  make([]uint, common.PuzzleSize),
					regionToDigitMap:        make(map[rune]uint),
				}
			}

			rows[uint(cell.Y)].UpdatedRegions[cell.X] = []rune(segment.Region)[0]
			rows[uint(cell.Y)].OriginalRegionsWithMask[cell.X] = []rune(segment.Region)[0]
		}
	}

	fmt.Printf("Loaded %d rows\n", len(rows))
}

func GenerateRow(y uint, mask []bool) *Row {
	if y >= uint(len(rows)) {
		return nil
	}

	row := Row{}

	row.OriginalRegionsWithMask = make([]rune, len(rows[y].OriginalRegionsWithMask))
	copy(row.OriginalRegionsWithMask, rows[y].OriginalRegionsWithMask)

	row.UpdatedRegions = make([]rune, common.PuzzleSize)
	row.regionToDigitMap = make(map[rune]uint)

	// apply mask
	for i := 0; i < common.PuzzleSize; i++ {
		if mask[i] {
			row.OriginalRegionsWithMask[i] = common.ShadedCell
			row.UpdatedRegions[i] = common.ShadedCell
		}
	}

	return &row
}

func (r *Row) DisplayRegions() {
	for _, symbol := range r.UpdatedRegions {
		fmt.Printf("%c", symbol)
	}
	fmt.Println()
}

func (r *Row) DisplayDigits() {
	for i, number := range r.Digits {
		if r.OriginalRegionsWithMask[i] == common.ShadedCell {
			fmt.Printf("%c", common.ShadedCell)
		} else {
			fmt.Printf("%d", number)
		}
	}
	fmt.Println()
}

func (r *Row) GetRegionToDigitMap() map[rune]uint {
	return r.regionToDigitMap
}

func (r *Row) ResetAssignments() {
	r.Digits = make([]uint, common.PuzzleSize)
	r.regionToDigitMap = make(map[rune]uint)
}

func (r *Row) AssignDigit(symbol rune, number uint) {
	for i := 0; i < len(r.UpdatedRegions); i++ {
		if r.UpdatedRegions[i] == symbol {
			r.Digits[i] = number
			r.regionToDigitMap[symbol] = number
		}
	}
}

func (r *Row) GetNumbers() []uint {
	var (
		numbers []uint
		Digits  []uint
	)

	sliceToNumber := func(slice []uint) uint64 {
		var result uint64
		for _, digit := range slice {
			result = result*10 + uint64(digit)
		}
		return result
	}

	for i := 0; i < len(r.UpdatedRegions); i++ {
		if r.UpdatedRegions[i] == common.ShadedCell {
			if len(Digits) != 0 {
				numbers = append(numbers, uint(sliceToNumber(Digits)))
			}
			Digits = []uint{}
			continue
		}

		Digits = append(Digits, r.Digits[i])
	}

	if len(Digits) != 0 {
		numbers = append(numbers, uint(sliceToNumber(Digits)))
	}

	return numbers
}

func (r *Row) GetRegionsToAssign(existingRegions map[rune]uint) []rune {
	var uniqueSymbols []rune

	for i := 0; i < len(r.UpdatedRegions); i++ {
		symbol := r.UpdatedRegions[i]

		if symbol == common.ShadedCell {
			continue
		}

		_, ok := existingRegions[symbol]
		if ok {
			continue
		}

		if !slices.Contains(uniqueSymbols, symbol) {
			uniqueSymbols = append(uniqueSymbols, symbol)
		}
	}

	return uniqueSymbols
}

func (r *Row) IsMaskCompatibleWith(mask []bool) bool {
	for i := 0; i < len(r.OriginalRegionsWithMask); i++ {
		if r.OriginalRegionsWithMask[i] == common.ShadedCell && mask[i] {
			return false
		}
	}

	return true
}
