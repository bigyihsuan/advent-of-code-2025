package util

import (
	"fmt"
	"io"
	"iter"
	"os"
	"strings"
)

// parser is a function that takes a line *WITH NEWLINE* and turns it into T
func LoadLines[T any](filename string, parser func(string) T) iter.Seq[T] {
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
