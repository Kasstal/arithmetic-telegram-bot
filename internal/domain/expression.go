package domain

import "fmt"

// ErrorType represents the type of error that can occur in the calculator.
type ErrorType string

// Expression represents a mathematical expression with its answer.
type Expression string

// Answer represents the answer to a mathematical expression.
type Answer float64

const (
	ErrorTypeInvalidExpression ErrorType = "invalid_expression"
	ErrorTypeDivisionByZero    ErrorType = "division_by_zero"
	ErrorTypeUnknown           ErrorType = "unknown_error"
)

type CalculatorError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *CalculatorError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func (e *CalculatorError) Is(target error) bool {
	if ce, ok := target.(*CalculatorError); ok {
		return e.Type == ce.Type
	}
	return false
}

func (e *CalculatorError) Unwrap() error {
	return e.Err
}
