package game

import (
	"fmt"
	"strings"
	"testing"

	"github.com/davidnotplay/gameoflife/matrix"
	"github.com/stretchr/testify/assert"
)

const min int = matrix.MINIMUM_SIZE
const disabled int = matrix.MATRIX_POINT_DISABLED
const enabled int = matrix.MATRIX_POINT_ENABLED

type adjTest struct {
	name             string
	position         [3]int
	adjacentsEnabled []Position
	countEnabled     int
}

// Make new game.
func TestMakeNewGame(t *testing.T) {
	assert := assert.New(t)
	g, err := New(min, min, []Position{})
	m := g.matrix

	assert.Equal(err, nil, "There an error when it made the game.")
	assert.Equal(m.GetWidth(), min, "Invalid width.")
	assert.Equal(m.GetHeight(), min, "Invalid height.")
}

//Make new game with some positions enabled intially
func TestMakeNewGamePosition(t *testing.T) {
	assert := assert.New(t)
	g, _ := New(min, min, []Position{{0, 0}, {1, 1}, {2, 2}})
	m := g.matrix

	r, _ := m.IsEnabled(0, 0)
	assert.Equal(r, true, "The position 0, 0 is disabled.")

	r, _ = m.IsEnabled(1, 1)
	assert.Equal(r, true, "The position 1, 1 is disabled.")

	r, _ = m.IsEnabled(2, 2)
	assert.Equal(r, true, "The position 2 2 is disabled.")
}

// Check error when it mades a game with invalid size.
func TestMakeNewGameSizeError(t *testing.T) {
	assert := assert.New(t)
	_, err := New(-1, -1, []Position{})
	assert.Equal(err, matrix.InvalidSizeError(-1, -1), "The error does not match.")
}

// Test error when use an invalid position.
func TestMakeNewGamePositionError(t *testing.T) {
	assert := assert.New(t)
	g, err := New(min, min, []Position{{min, min}})
	m := g.matrix

	assert.Equal(err, matrix.OutIndexError(m, min, min), "The error does not match.")

	g, err = New(min, min, []Position{{-1, -1}})
	m = g.matrix
	assert.Equal(err, matrix.OutIndexError(m, -1, -1), "The error does not match.")
}

func auxEnablePositions(g *Game, positions []Position) {
	for _, p := range positions {
		g.matrix.EnablePoint(p[0], p[1])
	}
}

func auxSetMessage(adt adjTest) string {
	strPositions := make([]string, len(adt.adjacentsEnabled))

	for i, p := range adt.adjacentsEnabled {
		strPositions[i] = fmt.Sprintf("\t- %dx%d\n", p[0], p[1])
	}

	message := "Test %s is invalid.\nPosition checked: (%dx%d).\nPositions enabled:\n%s"
	return fmt.Sprintf(
		message,
		adt.name,
		adt.position[0],
		adt.position[1],
		strings.Join(strPositions, ""))
}

// Test the function coundAdjacents.
func TestCoundAdjacentsFunc(t *testing.T) {
	assert := assert.New(t)
	g, _ := New(min, min, []Position{})

	// positions
	positions := []adjTest{
		{
			"test 1, left-top",
			[3]int{0, 0, disabled},
			[]Position{{2, 2}},
			0,
		},
		{
			"test 1, enabled, left-top",
			[3]int{0, 0, enabled},
			[]Position{{2, 2}},
			0,
		},
		{
			"test 2, left-top",
			[3]int{0, 0, disabled},
			[]Position{{1, 0}, {2, 2}},
			1,
		},
		{
			"test 3, left-top",
			[3]int{0, 0, disabled},
			[]Position{{1, 0}, {0, 1}, {2, 2}},
			2,
		},
		{
			"test 4, left-top",
			[3]int{0, 0, disabled},
			[]Position{{1, 0}, {1, 1}, {0, 1}, {2, 2}},
			3,
		},
		{
			"test 1, left-middle",
			[3]int{0, 5, disabled},
			[]Position{{9, 9}},
			0,
		},
		{
			"test 2, left-middle",
			[3]int{0, 5, enabled},
			[]Position{{0, 4}, {0, 6}, {9, 9}},
			2,
		},
		{
			"test 3, left-middle",
			[3]int{0, 5, enabled},
			[]Position{{0, 4}, {0, 6}, {1, 5}, {9, 9}},
			3,
		},
		{
			"test 1, left-bottom",
			[3]int{0, 9, disabled},
			[]Position{{9, 9}},
			0,
		},
		{
			"test 2, left-bottom",
			[3]int{0, 9, disabled},
			[]Position{{0, 8}, {9, 9}},
			1,
		},
		{
			"test 3, left-bottom",
			[3]int{0, 9, disabled},
			[]Position{{0, 8}, {1, 9}, {9, 9}},
			2,
		},
		{
			"test 1, center-top",
			[3]int{5, 0, disabled},
			[]Position{{9, 9}},
			0,
		},
		{
			"test 2, center-top",
			[3]int{5, 0, disabled},
			[]Position{{5, 1}, {4, 0}, {9, 9}},
			2,
		},
		{
			"test 3, center-top",
			[3]int{5, 0, disabled},
			[]Position{{6, 0}, {4, 0}, {6, 1}, {9, 9}},
			3,
		},
		{
			"test 1, center-middle",
			[3]int{5, 5, disabled},
			[]Position{{9, 9}},
			0,
		},
		{
			"test 2, center-middle",
			[3]int{5, 5, disabled},
			[]Position{{4, 4}, {6, 6}, {9, 9}},
			2,
		},
		{
			"test 3, center-middle",
			[3]int{5, 5, enabled},
			[]Position{{4, 4}, {6, 6}, {9, 9}},
			2,
		},
		{
			"test 4, center-middle",
			[3]int{5, 5, enabled},
			[]Position{{4, 4}, {6, 6}, {5, 6}, {4, 5}, {9, 9}},
			4,
		},
		{
			"test 1, right-middle",
			[3]int{9, 5, enabled},
			[]Position{{9, 9}},
			0,
		},
		{
			"test 2, right-middle",
			[3]int{9, 5, enabled},
			[]Position{{9, 4}, {9, 6}, {8, 5}, {9, 9}},
			3,
		},
		{
			"test 1, left-bottom",
			[3]int{0, 9, enabled},
			[]Position{{9, 9}},
			0,
		},
		{
			"test 2, left-bottom",
			[3]int{0, 9, enabled},
			[]Position{{0, 8}, {1, 9}, {9, 9}},
			2,
		},
		{
			"test 1, center-bottom",
			[3]int{5, 9, enabled},
			[]Position{{0, 8}, {1, 9}, {9, 9}},
			0,
		},
		{
			"test 2, center-bottom",
			[3]int{5, 9, enabled},
			[]Position{{4, 9}, {5, 8}, {9, 9}},
			2,
		},
		{
			"test 2, center-right",
			[3]int{9, 9, enabled},
			[]Position{{4, 9}, {5, 8}, {9, 9}},
			0,
		},
		{
			"test 2, center-right",
			[3]int{9, 9, enabled},
			[]Position{{8, 9}, {9, 8}, {9, 9}},
			2,
		},
	}

	for _, adt := range positions {
		g.matrix.Reset()
		auxEnablePositions(g, adt.adjacentsEnabled)
		x, y := adt.position[0], adt.position[1]

		if adt.position[2] == enabled {
			g.matrix.EnablePoint(x, y)
		}

		assert.Equal(g.countAdjacents(x, y), adt.countEnabled, auxSetMessage(adt))
	}
}

func TestGetMatrix(t *testing.T) {
	assert := assert.New(t)
	g, _ := New(min, min, []Position{{0, 0}, {1, 1}, {2, 2}})

	m := g.GetMatrix()

	enabled, _ := m.IsEnabled(0, 0)
	assert.Equal(enabled, true, "The point 0x0 is disabled.")

	enabled, _ = m.IsEnabled(1, 1)
	assert.Equal(enabled, true, "The point 1x1 is disabled.")

	enabled, _ = m.IsEnabled(2, 2)
	assert.Equal(enabled, true, "The point 2x2 is disabled.")

	enabled, _ = m.IsEnabled(3, 3)
	assert.Equal(enabled, false, "The point 3x3 is enabled.")
}

// When a live cell is adjacent to 2 or 3 live cells this lives
func TestRule1InFuncRule(t *testing.T) {
	assert := assert.New(t)
	var err error = nil

	// 2 live cell adjacents.
	g, _ := New(min, min, []Position{{0, 0}, {1, 1}, {2, 2}})

	adj := g.countAdjacents(1, 1)
	g.rules(true, adj, 1, 1, &err)
	enabled, _ := g.GetMatrix().IsEnabled(1, 1)
	assert.Equal(err, nil, "There is an error.j")
	assert.Equal(enabled, true, "The point 1x1 is disabled.")

	// 3 cells adjacents
	g, _ = New(min, min, []Position{{0, 0}, {1, 1}, {2, 1}, {2, 2}})
	adj = g.countAdjacents(1, 1)
	g.rules(true, adj, 1, 1, &err)
	enabled, _ = g.GetMatrix().IsEnabled(1, 1)
	assert.Equal(err, nil, "There is an error.j")
	assert.Equal(enabled, true, "The point 1x1 is disabled.")
}

// When a live cell is adjacents to 1 or more than 3 live cells this dead.
func TestRule2InFuncRule(t *testing.T) {
	assert := assert.New(t)
	var err error = nil

	// 1 live cell adjacents.
	g, _ := New(min, min, []Position{{0, 0}, {1, 1}})
	adj := g.countAdjacents(1, 1)
	g.rules(true, adj, 1, 1, &err)
	enabled, _ := g.GetMatrix().IsEnabled(1, 1)
	assert.Equal(err, nil, "There is an error.")
	assert.Equal(enabled, false, "The point is enabled.")

	// 4 or more cells lives.
	g, _ = New(min, min, []Position{{0, 0}, {1, 1}, {1, 0}, {0, 1}, {1, 2}})
	adj = g.countAdjacents(1, 1)
	g.rules(true, adj, 1, 1, &err)
	enabled, _ = g.GetMatrix().IsEnabled(1, 1)
	assert.Equal(err, nil, "There is an error.")
	assert.Equal(enabled, false, "The point is enabled.")
}

// When a dead cell is adjacents to number different of 3 live cells it continues dead.
func TestRule3InFuncRule(t *testing.T) {
	assert := assert.New(t)
	var err error = nil

	// 1 live cell adjacents.
	g, _ := New(min, min, []Position{{0, 0}})
	adj := g.countAdjacents(1, 1)
	g.rules(false, adj, 1, 1, &err)
	enabled, _ := g.GetMatrix().IsEnabled(1, 1)
	assert.Equal(err, nil, "There is an error.")
	assert.Equal(enabled, false, "The point is enabled.")

	// 4 live cells adjacents.
	g, _ = New(min, min, []Position{{0, 0}, {0, 1}, {0, 2}, {1, 0}})
	adj = g.countAdjacents(1, 1)
	g.rules(false, adj, 1, 1, &err)
	enabled, _ = g.GetMatrix().IsEnabled(1, 1)
	assert.Equal(err, nil, "There is an error.")
	assert.Equal(enabled, false, "The point is enabled.")
}

// When a dead cell is adjacents to 3 live cells it gets born.
func TestRule4InFuncRule(t *testing.T) {
	assert := assert.New(t)
	var err error = nil

	// 3 live cells adjacents.
	g, _ := New(min, min, []Position{{0, 0}, {0, 1}, {0, 2}})
	adj := g.countAdjacents(1, 1)
	g.rules(false, adj, 1, 1, &err)
	enabled, _ := g.GetMatrix().IsEnabled(1, 1)
	assert.Equal(err, nil, "There is an error.")
	assert.Equal(enabled, true, "The point is disabled.")
}

// Test when the function `game.rules` generate an error.
func TestErrorInFuncrule(t *testing.T) {
	var err error = nil
	assert := assert.New(t)

	g, _ := New(min, min, []Position{})
	g.rules(true, 0, 11, 11, &err)
	assert.Equal(err, matrix.OutIndexError(g.matrix, 11, 11), "The error does not match.")
}

// Testing the function `game.Cycle`
func TestCycleFunc(t *testing.T) {
	g, _ := New(min, min, []Position{{1, 1}, {2, 1}, {3, 1}})
	assert := assert.New(t)
	message := "The position %dx%d is %s"
	pointStr := map[bool]string{true: "enabled", false: "disable"}
	positions := map[Position]bool{{2, 0}: true, {2, 1}: true, {2, 2}: true}

	//  run cycle
	err := g.Cycle()

	// no error
	assert.Equal(err, nil, "There is an error.")

	// Check the point status.
	for i := 0; i < g.matrix.GetWidth(); i++ {
		for j := 0; j < g.matrix.GetHeight(); j++ {
			enabled, _ := g.matrix.IsEnabled(i, j)
			_, ok := positions[Position{i, j}]
			assert.Equal(enabled, ok, fmt.Sprintf(message, i, j, pointStr[enabled]))

		}
	}

	// run new cycle
	positions = map[Position]bool{{1, 1}: true, {2, 1}: true, {3, 1}: true}
	err = g.Cycle()

	// no error
	assert.Equal(err, nil, "There is an error.")

	// Check the point status.
	for i := 0; i < g.matrix.GetWidth(); i++ {
		for j := 0; j < g.matrix.GetHeight(); j++ {
			enabled, _ := g.matrix.IsEnabled(i, j)
			_, ok := positions[Position{i, j}]
			assert.Equal(enabled, ok, fmt.Sprintf(message, i, j, pointStr[enabled]))

		}
	}
}
