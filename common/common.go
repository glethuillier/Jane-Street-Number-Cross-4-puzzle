package common

import (
	"path/filepath"
	"runtime"
)

const (
	PuzzleSize = 11
	ShadedCell = '*'
)

type Coordinates struct {
	X int
	Y int
}

type Cell struct {
	Coordinates
	OriginalRegion rune
	UpdatedRegion  rune
}

func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
