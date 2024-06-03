package solver

import (
	"fmt"
	"jane/common"
	"jane/row"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getOriginalRegions(y uint, candidate *Candidate, current *row.Row) []string {
	regions := []string{
		"(Original regions with shaded cells)",
	}

	for i := uint(0); i < y; i++ {
		regions = append(
			regions,
			string(candidate.Rows[i].OriginalRegionsWithMask),
		)
	}
	regions = append(regions,
		string(current.OriginalRegionsWithMask),
	)

	return regions
}

func getUpdatedRegions(y uint, candidate *Candidate, current *row.Row) []string {
	regions := []string{
		"\n(Updated regions)",
	}

	for i := uint(0); i < y; i++ {
		regions = append(
			regions,
			string(candidate.Rows[i].UpdatedRegions),
		)
	}
	regions = append(regions,
		string(current.UpdatedRegions),
	)

	return regions
}

func getDigits(y uint, candidate *Candidate, current *row.Row) []string {
	digits := []string{
		"\n(Digits)",
	}

	var d string
	for i := uint(0); i < y; i++ {
		d = ""
		for j, s := range candidate.Rows[i].OriginalRegionsWithMask {
			if s == common.ShadedCell {
				d += string(common.ShadedCell)
			} else {
				d += fmt.Sprintf("%d", candidate.Rows[i].Digits[j])
			}
		}
		digits = append(digits, d)
	}

	d = ""
	for j, s := range current.OriginalRegionsWithMask {
		if s == common.ShadedCell {
			d += string(common.ShadedCell)
		} else {
			d += fmt.Sprintf("%d", current.Digits[j])
		}
	}
	digits = append(digits, d)

	return digits
}

func getNumbers(y uint, candidate *Candidate, current *row.Row) []string {
	numbers := []string{
		"\n(Numbers)",
	}

	sliceToString := func(slice []uint) []string {
		strSlice := make([]string, len(slice))

		for i, val := range slice {
			strSlice[i] = strconv.FormatUint(uint64(val), 10)
		}

		return strSlice
	}

	for i := uint(0); i < y; i++ {
		numbers = append(numbers,
			strings.Join(sliceToString(candidate.Rows[i].GetNumbers()), ", "),
		)
	}
	numbers = append(numbers,
		strings.Join(sliceToString(current.GetNumbers()), ", "),
	)

	return numbers
}

func SaveCandidate(y uint, candidate *Candidate, current *row.Row) {
	var ls []string

	ls = append(ls, getOriginalRegions(y, candidate, current)...)
	ls = append(ls, getUpdatedRegions(y, candidate, current)...)
	ls = append(ls, getDigits(y, candidate, current)...)
	ls = append(ls, getNumbers(y, candidate, current)...)
	ls = append(ls, "\n--------------------\n\n")

	filename := fmt.Sprintf("%d_", y+1)
	if y < 1 {
		filename += "row"
	} else {
		filename += "rows"
	}

	file, err := os.OpenFile(
		filepath.Clean(filepath.Join(
			common.GetBasePath(),
			"..",
			directoryPath,
			fmt.Sprintf("%s.txt", filename),
		)),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600,
	)
	if err != nil {
		panic(err)
	}

	if _, err := file.Write([]byte(strings.Join(ls, "\n"))); err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}
