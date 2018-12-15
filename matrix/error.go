package matrix

import (
	"fmt"
)

type outIndexError struct {
	size     [2]int
	position [2]int
}

func (e *outIndexError) Error() string {
	message := "Position (%d, %d) not found in matrix (%d, %d)."
	return fmt.Sprintf(
		message,
		e.position[0],
		e.position[1],
		e.size[0],
		e.size[1])
}

type invalidSizeError [2]int

func (self *invalidSizeError) Error() string {
	message := "The matrix size (%dx%d) is invalid. Minimum (%dx%d)"
	return fmt.Sprintf(message, self[0], self[1], MINIMUM_SIZE, MINIMUM_SIZE)
}

func OutIndexError(m *Matrix, x, y int) error {
	err := outIndexError{}

	err.size[0] = m.GetWidth()
	err.size[1] = m.GetHeight()

	err.position[0] = x
	err.position[1] = y
	return &err
}


func InvalidSizeError(width, height int) error {
	return &invalidSizeError{width, height}
}
