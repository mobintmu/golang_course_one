package pointer

import "errors"

type Point struct {
	x, y int
}

type Sample struct {
	input  string
	output string
}

func Change(x, y int) {
	x++
	y++
}

func ChangePointer(x, y *int) {
	*x++
	*y++
}

func NewPoint(x, y int) (*Point, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("x and y must be >= 0")
	}
	return &Point{x, y}, nil

}
