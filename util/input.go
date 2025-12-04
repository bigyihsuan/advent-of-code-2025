package util

import (
	"aoc2025/util/grid"
	"aoc2025/util/strpipe"
	"fmt"
	"io"
	"iter"
	"os"
	"strings"
)

type Parser[T any] func(string) T

// parser is a function that takes a line *WITH NEWLINE* and turns it into T
func LoadLines[T any](filename string, parser Parser[T]) iter.Seq[T] {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(fmt.Errorf("loading lines: getting file: %w", err))
	}
	defer file.Close()
	bts, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("loading lines: reading file: %w", err))
	}
	lines := strings.Lines(string(bts))

	return func(yield func(T) bool) {
		for line := range lines {
			if !yield(parser(line)) {
				return
			}
		}
	}
}

func LoadLineOfElements[T any](filename string, parser Parser[T]) iter.Seq[T] {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(fmt.Errorf("loading line: getting file: %w", err))
	}
	defer file.Close()
	bts, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("loading line: reading file: %w", err))
	}
	sp := strpipe.New(string(bts))

	return func(yield func(T) bool) {
		for element := range sp.Trim().SplitCommas() {
			if !yield(parser(element.String())) {
				return
			}
		}
	}
}

func LoadStringGrid(filename string) grid.Grid {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(fmt.Errorf("loading grid: getting file: %w", err))
	}
	defer file.Close()
	bts, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("loading grid: reading file: %w", err))
	}

	return grid.NewFromStringGrid(string(bts))
}
