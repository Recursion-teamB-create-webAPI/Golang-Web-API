package testErros

import "fmt"

type InitMockDbError struct {}

func (e *InitMockDbError) Error() string {
	return fmt.Sprintln("Failed to init db mock")
}

func NewInitMockDbError() *InitMockDbError {
	return &InitMockDbError{}
}

