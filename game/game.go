package game

import (
	"github.com/davidnotplay/gameoflife/matrix"
)

type Game struct {
	matrix *matrix.Matrix
	cycles uint
}

type Position [2]int

// Make new game with a matrix of size `width`x`height`
// Param `position` allow define the initial cells enabled in the matrix.
// returns the type Game or an error.
func New(width, height int, positions []Position) (*Game, error) {
	var err error = nil
	m, err := matrix.New(width, height)

	// Check if the matrix has an error.
	if err != nil {
		return nil, err
	}

	g := &Game{m, 0}

	// Enable the initial positions.
	for _, position := range positions {
		err = m.EnablePoint(position[0], position[1])
	}

	if err != nil {
		// There are errors in the initial positions.
		return g, err
	}

	return g, nil
}

// Returns the numbers of point enabled of the matrix `self.matrix` around of the point `x`, `y`
func (self *Game) countAdjacents(x, y int) int {
	var xmin, xmax, ymin, ymax, count int
	width, height := self.matrix.GetSize()
	increment := map[bool]int{false: 0, true: 1}

	// Define the adjacent zone to the point x,y
	if xmin = x; x > 0 {
		xmin = x - 1
	}

	if xmax = x; x < width-1 {
		xmax = x + 1
	}

	if ymin = y; y > 0 {
		ymin = y - 1
	}

	if ymax = y; y < height-1 {
		ymax = y + 1
	}

	count = 0

	// Count adjacents.
	for i := xmin; i <= xmax; i++ {
		for j := ymin; j <= ymax; j++ {
			enabled, _ := self.matrix.IsEnabled(i, j)
			count += increment[enabled]
		}
	}

	// Decrement the point x, y whether it is enabled.
	enabled, _ := self.matrix.IsEnabled(x, y)
	count -= increment[enabled]

	return count
}

// Returns the mtrix game.
func (self *Game) GetMatrix() *matrix.Matrix {
	return self.matrix
}

// Using the game rules, modify status of the point `x`, `y`, depending of the param `enabled`,
// that indicates if the point is enabled, and the  param `adj` that is the number of enabled
// points adjacents. The **out** param `error` is used for returns the errors.
func (self *Game) rules(enabled bool, adj, x, y int, err *error) {
	if *err != nil {
		// There is an error from parent.
		return
	}

	// Apply the game rules depending of `adj` and `enabled`
	// When point is enabled
	if enabled {
		if adj != 2 && adj != 3 {
			*err = self.matrix.DisablePoint(x, y)
			return
		}
	} else if adj == 3 {
		*err = self.matrix.EnablePoint(x, y)
		return
	}

	*err = nil
}

// The func run all points in matrix, apply the rules in they
// and generate a new state of the matrix.
func (self *Game) Cycle() (err error) {
	err = nil
	width, height := self.matrix.GetWidth(), self.matrix.GetHeight()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			enabled, err := self.matrix.IsEnabled(i, j)

			if err != nil {
				return err
			}

			adj := self.countAdjacents(i, j)
			// Using deffer check the rules and apply the new state of the point.
			defer self.rules(enabled, adj, i, j, &err)
		}
	}

	self.cycles++
	return
}

// Get the number of cycles.
func (self *Game) GetCyclesNum() uint {
	return self.cycles
}
