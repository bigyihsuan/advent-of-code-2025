package day01

import (
	rt "aoc2025/day01/rotation"
	"aoc2025/util"
	"fmt"
	"iter"
)

func Day(filename string, part int) {
	rs := util.LoadLines(filename, rt.Parse)
	switch part {
	case 1:
		fmt.Println(TurnLock(rs))
	case 2:
		fmt.Println(TurnLock2(rs))
	}
}

func TurnLock(rs iter.Seq[rt.Rotation]) int {
	zeroes := 0
	current := 50
	for r := range rs {
		if r.Direction == rt.Left {
			current = (current - r.Clicks) % 100
		} else {
			current = (current + r.Clicks) % 100
		}
		if current == 0 {
			zeroes++
		}
	}
	return zeroes
}

func TurnLock2(rs iter.Seq[rt.Rotation]) int {
	zeroes := 0
	current := 50
	// fmt.Println(current, zeroes)
	for r := range rs {
		clicks := r.Clicks % 100
		zeroes += r.Clicks / 100
		if r.Direction == rt.Left {
			c := current - clicks
			if current != 0 && c < 0 {
				// fmt.Println(c)
				zeroes += 1
			}
			current = (100 + c) % 100
		} else {
			c := current + clicks
			if current != 0 && c > 100 {
				// fmt.Println(c)
				zeroes += 1
			}
			current = c % 100
		}
		if current == 0 {
			zeroes++
		}
		// fmt.Println(r, current, zeroes)
	}
	return zeroes
}
