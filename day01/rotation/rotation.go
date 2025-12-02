package rotation

import (
	"fmt"
	"strconv"
	"strings"
)

type Direction bool

const (
	Left  Direction = true  // L
	Right           = false // R
)

func (d Direction) String() string {
	if d == Left {
		return "L"
	} else {
		return "R"
	}
}

type Rotation struct {
	Direction
	Clicks int
}

func Parse(line string) Rotation {
	var r Rotation
	line = strings.TrimSpace(line)
	dir, clicks := line[0], line[1:]
	if dir == 'L' {
		r.Direction = Left
	} else {
		r.Direction = Right
	}

	c, err := strconv.ParseInt(clicks, 10, 64)
	if err != nil {
		panic(fmt.Errorf("parsing rotation: %w", err))
	}
	r.Clicks = int(c)
	return r
}

func (r Rotation) String() string {
	return fmt.Sprintf("%s%d", r.Direction.String(), r.Clicks)
}
