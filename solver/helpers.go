package solver

import (
	"fmt"
	"jane/common"
	"jane/row"
	"os"
	"path/filepath"
	"sync/atomic"
)

const directoryPath = "./candidates/"

var counter int64

func init() {
	// ensure that the directory where the candidates
	// will be stored exists and is empty
	err := ensureFolder(directoryPath)
	if err != nil {
		panic(err)
	}

	atomic.StoreInt64(&counter, -1)
}

// counter

func isGreaterThanCounter(newValue int64) bool {
	return atomic.LoadInt64(&counter) < newValue
}

func setCounter(newValue int64) {
	atomic.StoreInt64(&counter, newValue)
}

// candidates directory

func ensureFolder(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0750)
		if err != nil {
			panic(err)
		}
	} else {
		err = cleanDirectory(path)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func cleanDirectory(path string) error {
	dir, err := os.Open(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("failed to open directory: %w", err)
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	err = dir.Close()
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(path, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// display candidates

func displayRows(rowID uint, candidate *Candidate, currentRow *row.Row) {
	switch i := rowID; {
	case i < 1:
		fmt.Printf("--- %d row ---\n", rowID+1)
	case i < common.PuzzleSize-1:
		fmt.Printf("--- %d rows ---\n", rowID+1)
	default:
		fmt.Printf("--- SOLUTION ---\n")
	}

	fmt.Println("(regions)")
	for i := range rowID {
		candidate.Rows[i].DisplayRegions()
	}
	currentRow.DisplayRegions()

	fmt.Println()

	fmt.Println("(digits)")
	for i := range rowID {
		candidate.Rows[i].DisplayDigits()
	}
	currentRow.DisplayDigits()

	fmt.Println()
}
