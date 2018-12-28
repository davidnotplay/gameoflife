package matrix

import (
	"fmt"
)

// Value of the disabled points in the matrix.
const MATRIX_POINT_DISABLED int = 0

// Value of the enabled points in the matrix.
const MATRIX_POINT_ENABLED int = 1

// Matrix minum size.
const MINIMUM_SIZE int = 10

// Matrix base struct.
type Matrix struct {
	// Multi-Slice with the data.
	matrix [][]int

	// Slice horizontal size.
	width int

	// Slice vertical size.
	height int

	// points enabled.
	enabled int
}

// Check if the `x`, `y` position are inside range of the matrix.
// Returns nil if all is correct or an error whether the poisition is invalid.
func checkRange(m *Matrix, x, y int) error {
	w, h := m.width, m.height
	if w <= x || h <= y || x < 0 || y < 0 {
		return OutIndexError(m, x, y)
	}

	return nil
}

// Creates an array of size `width`,`height`
// where all his cells have the value `MATRIX_POINT_DISABLED`
func createEmptyMatrixArray(width int, height int) (m [][]int) {
	m = make([][]int, width)

	for i := 0; i < width; i++ {
		m[i] = make([]int, height)
		for j := 0; j < height; j++ {
			m[i][j] = MATRIX_POINT_DISABLED
		}
	}
	return
}

// Creates and returns a new instance of the struct `Matrix`.
// `width`and `height` params are used to define the matrix size.
// It returns and error whether width or height are negative numbers.
func New(width, height int) (*Matrix, error) {
	if width < MINIMUM_SIZE || height < MINIMUM_SIZE {
		err := InvalidSizeError(width, height)
		return nil, err
	}

	m := &Matrix{createEmptyMatrixArray(width, height), width, height, 0}
	return m, nil
}

// Enable the point of the position `x`, `y` of the matrix stored in `self`.
// Whether the position is invalid returns an error.
func (self *Matrix) EnablePoint(x, y int) (e error) {
	e = checkRange(self, x, y)
	if e != nil {
		return e
	}

	if self.matrix[x][y] != MATRIX_POINT_ENABLED {
		self.matrix[x][y] = MATRIX_POINT_ENABLED
		self.enabled++
	}

	return nil
}

// Disable the point of the position `x`, `y` of the matrix stored in `self`.
// Whether the position is invalid returns an error.
func (self *Matrix) DisablePoint(x, y int) (e error) {
	e = checkRange(self, x, y)

	if e != nil {
		return e
	}

	if self.matrix[x][y] != MATRIX_POINT_DISABLED {
		self.matrix[x][y] = MATRIX_POINT_DISABLED
		self.enabled--
	}

	return nil
}

// Checks if the point of the position `x`, `y` of the matrix stored in `self` is enabled.
// Whether the position is invalid returns error.
func (self *Matrix) IsEnabled(x, y int) (bool, error) {
	e := checkRange(self, x, y)
	if e == nil {
		return self.matrix[x][y] == MATRIX_POINT_ENABLED, nil
	}

	return false, e
}

// Disable all points in the matrix.
func (self *Matrix) Reset() {
	for i := 0; i < self.width; i++ {
		for j := 0; j < self.height; j++ {
			self.matrix[i][j] = MATRIX_POINT_DISABLED
		}
	}
}

// Returns the value of the position `x`, `y` of the matrix stored in `self`.
// Whether position is invalid returns an error as second element.
func (self *Matrix) GetPoint(x, y int) (int, error) {
	e := checkRange(self, x, y)
	if e == nil {
		return self.matrix[x][y], nil
	}

	return 0, e
}

// Returns the *width* of the matrix stored in `self`
func (self *Matrix) GetWidth() int {
	return self.width
}

// Returns the *height* of the matrix stored in `self`
func (self *Matrix) GetHeight() int {
	return self.height
}

// returns the width and height of the matrix.
func (self *Matrix) GetSize() (int, int) {
	return self.width, self.height
}

// Get the points enabled.
func (self *Matrix) GetPointsEnabled() int {
	return self.enabled
}

func (matrix Matrix) String() string {
	msg := "Matrix (%dx%d)"
	return fmt.Sprintf(msg, matrix.width, matrix.height)
}
