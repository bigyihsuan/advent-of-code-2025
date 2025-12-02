package main

import (
	"aoc2025/day01"
	"path"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Day    int  `short:"d" long:"day"`
	Part   int  `short:"p" long:"part" default:"1"`
	Sample bool `short:"s" long:"sample"`
}

type DayFunc = func(filename string, part int)

var Days = map[int]struct {
	a DayFunc
	f string
}{
	1: {a: day01.Day01, f: "day01"},
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	d := Days[opts.Day]
	filename := d.f
	if opts.Sample {
		filename = path.Join("sample", filename)
	} else {
		filename = path.Join("input", filename)
	}

	d.a(filename, opts.Part)
}
