package main

import (
	"aoc2025/day01"
	"aoc2025/day02"
	"aoc2025/day03"
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
	2: {a: day02.Day02, f: "day02"},
	3: {a: day03.Day03, f: "day03"},
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
		filename = path.Join("sample", filename+".txt")
	} else {
		filename = path.Join("input", filename+".txt")
	}

	d.a(filename, opts.Part)
}
