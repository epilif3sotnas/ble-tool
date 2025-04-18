package cli


import (
	// std
	"fmt"
)


type HexStringTooLong struct {
	Field string
	FieldSize int
	Size uint16
}

func NewHexStringTooLong(field string, fieldSize int, size uint16) *HexStringTooLong {
	return &HexStringTooLong{
		Field: field,
		FieldSize: fieldSize,
		Size: size,
	}
}

func (e *HexStringTooLong) Error() string {
	return fmt.Sprintf("Hex String inserted (field %s) does not match the size required. Size inserted %d and requires %d.", e.Field, e.FieldSize, e.Size)
}


type HexStringTooShort struct {
	Field string
	FieldSize int
	Size uint16
}

func NewHexStringTooShort(field string, fieldSize int, size uint16) *HexStringTooShort {
	return &HexStringTooShort{
		Field: field,
		FieldSize: fieldSize,
		Size: size,
	}
}

func (e *HexStringTooShort) Error() string {
	return fmt.Sprintf("Hex String inserted (field %s) does not match the size required. Size inserted %d and requires %d.", e.Field, e.FieldSize, e.Size)
}


type CommandNotSupported struct {
	Command string
}
func NewCommandNotSupported(command string) *CommandNotSupported {
	return &CommandNotSupported{
		Command: command,
	}
}

func (e *CommandNotSupported) Error() string {
	return fmt.Sprintf("Command %s not supported.", e.Command)
}