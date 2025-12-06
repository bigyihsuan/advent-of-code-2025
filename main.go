package main

import (
	"aoc2025/day01"
	"aoc2025/day02"
	"aoc2025/day03"
	"aoc2025/day04"
	"aoc2025/day05"
	"aoc2025/day06"
	"path"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Day      int    `short:"d" long:"day"`
	Part     int    `short:"p" long:"part" default:"1"`
	Sample   bool   `short:"s" long:"sample"`
	Filename string `short:"f" long:"filename" default:""`
}

type DayFunc = func(filename string, part int)

var Days = map[int]struct {
	a DayFunc
	f string
}{
	1: {a: day01.Day, f: "day01"},
	2: {a: day02.Day, f: "day02"},
	3: {a: day03.Day, f: "day03"},
	4: {a: day04.Day, f: "day04"},
	5: {a: day05.Day, f: "day05"},
	6: {a: day06.Day, f: "day06"},
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
	} else if opts.Filename != "" {
		filename = opts.Filename
	} else {
		filename = path.Join("input", filename+".txt")
	}

	d.a(filename, opts.Part)
}
