package matrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const min = MINIMUM_SIZE

// Test the matrix size when it is made.
func TestNewMatrix(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	// test matrix width
	assert.Equal(len(m.matrix), min, "The matrix width is different.")

	for _, col := range m.matrix {
		// test the matrix height.
		assert.Equal(len(col), min, "The matrix height is different.")
	}
}

// Test if all values of the matrix slice are equal to matrix.MATRIX_POINT_DISABLED
func TestNewMatrixSliceSize(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	for i := 0; i < min; i++ {
		for j := 0; j < min; j++ {
			assert.Equal(
				m.matrix[i][j],
				MATRIX_POINT_DISABLED,
				fmt.Sprintf("The matrix point %d,%d is enabled.", i, j))
		}
	}
}

// Check if it throw an error when the size are invalid.
func TestNewMatrixInvalidSize(t *testing.T) {
	assert := assert.New(t)
	_, err := New(-1, -1)
	assert.Equal(err, InvalidSizeError(-1, -1), "The error does not match")
}

// Test Method GetPoint
func TestGetPoint(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	m.matrix[1][1] = MATRIX_POINT_ENABLED

	value, err := m.GetPoint(1, 1)
	assert.Equal(err, nil, "The error is not nil.")
	assert.Equal(value, MATRIX_POINT_ENABLED, "The point 1, 1 is not enabled.")
}

// Test error when pass an invalid position to GetPoint
func TestGetPointError(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	value, err := m.GetPoint(min, min)
	assert.Equal(value, 0, "The value is not zero.")
	assert.Equal(err, OutIndexError(m, min, min), "The error does not match")

	// negative numbers
	value, err = m.GetPoint(-1, -1)
	assert.Equal(value, 0, "The value is not zero.")
	assert.Equal(err, OutIndexError(m, -1, -1), "The error does not match")
}

// Test the Function EnablePoint
func TestEnablePoint(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	err := m.EnablePoint(0, 0)
	assert.Equal(err, nil, "The function does not return nil.")
	assert.Equal(m.matrix[0][0], MATRIX_POINT_ENABLED)
}

// Test error when pass an invalid position to EnablePoint
func TestEnablePointError(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	err := m.EnablePoint(min, min)
	assert.Equal(err, OutIndexError(m, min, min), "The error does not match.")

	err = m.EnablePoint(-1, -1)
	assert.Equal(err, OutIndexError(m, -1, -1), "The error does not match.") }
// Test the function DisablePoint
func TestDisablePoint(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	m.EnablePoint(0, 0)
	assert.Equal(m.matrix[0][0], MATRIX_POINT_ENABLED, "The point is not enabled.")

	m.DisablePoint(0, 0)
	assert.Equal(m.matrix[0][0], MATRIX_POINT_DISABLED, "The point is not disabled.")
}

// Test the error when pass an invalid position to DisablePoint.
func TestDisablePointError(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	err := m.DisablePoint(min, min)
	assert.Equal(err, OutIndexError(m, min, min), "The error does not match.")

	err = m.DisablePoint(-1, -1)
	assert.Equal(err, OutIndexError(m, -1, -1), "The error does not match.")
}

// Test the function EnablePoint.
func TestIsEnabledPoint(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	m.EnablePoint(1, 1)
	value, _ := m.IsEnabled(1, 1)
	assert.Equal(value, true, "The point is disabled.")
}

// Test the error when pass an invalid position to IsEnable function.
func TestIsEnablePointError(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	value, err := m.IsEnabled(min, min)
	assert.Equal(value, false, "The value is not false.")
	assert.Equal(err, OutIndexError(m, min, min), "The error does not match.")

	value, err = m.IsEnabled(-1, -1)
	assert.Equal(value, false, "The value is not false.")
	assert.Equal(err, OutIndexError(m, -1, -1), "The error does not match.")
}

// Test the function Reset
func TestResetMatrix(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)
	msg := "The point %d,%d is enabled"

	for i := 0; i < m.width; i++ {
		m.EnablePoint(i, i)
	}

	m.Reset()
	for i := 0; i < min; i++ {
		for j := 0; j < min; j++ {
			enabled, _ := m.IsEnabled(i, j)
			assert.Equal(enabled, false, fmt.Sprintf(msg, i, j))
		}
	}
}


// Test the functio GetWidth
func TestGetWidth(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)
	assert.Equal(m.GetWidth(), min)
}

// Test the function GetHeight
func TestGetHeight(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)
	assert.Equal(m.GetHeight(), min)
}

// Test the function GetSize
func TestGetSize(t *testing.T) {
	assert := assert.New(t)
	m, _ := New(min, min)

	width, height := m.GetSize()
	assert.Equal(width, min, "Invalid width.")
	assert.Equal(height, min, "Invalid height.")
}
