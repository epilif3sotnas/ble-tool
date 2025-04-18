package cli

import (
	// std
	"fmt"
	"testing"

	// external
	"github.com/stretchr/testify/assert"
)


func TestErrorHexStringTooLongWhenDataValid(t *testing.T) {
	field := "testField"
	fieldSize := 1
	size := uint16(3)


	expected := fmt.Sprintf("Hex String inserted (field %s) does not match the size required. Size inserted %d and requires %d.", field, fieldSize, size)
	actual := NewHexStringTooLong(field, fieldSize, size).Error()


	assert.Equalf(t, expected, actual, "Verification of error message for error HexStringTooLong")
}

func TestErrorHexStringTooShorthenDataValid(t *testing.T) {
	field := "testField"
	fieldSize := 10
	size := uint16(3)


	expected := fmt.Sprintf("Hex String inserted (field %s) does not match the size required. Size inserted %d and requires %d.", field, fieldSize, size)
	actual := NewHexStringTooLong(field, fieldSize, size).Error()


	assert.Equalf(t, expected, actual, "Verification of error message for error HexStringTooShort")
}