package tests

import (
	"arithmetic-telegram-bot/internal/domain"
	"arithmetic-telegram-bot/internal/infrastructure/evaluator"
	"log"
	"testing"
)

func TestEvaluate_ValidExpression(t *testing.T) {
	// Arrange
	evaluator := evaluator.NewGoEvaluator()

	testCases := []struct {
		expression string
		expected   float64
	}{
		{"2 + 2", 4},
		{"10 - 5", 5},
		{"3 * 4", 12},
		{"10 / 2", 5},
		{"(5 + 5) * 2", 20},
		{"10 + (5 * 2)", 20},
	}

	for _, tc := range testCases {
		// Act
		result, err := evaluator.Evaluate(domain.Expression(tc.expression))

		// Assert
		if err != nil {
			t.Errorf("Expected no error for expression '%s', got %v", tc.expression, err)
		}

		if float64(result) != tc.expected {
			t.Errorf("For expression '%s', expected result to be %f, got %f", tc.expression, tc.expected, float64(result))
		}
	}
}

func TestEvaluate_InvalidExpression(t *testing.T) {
	// Arrange
	evaluator := evaluator.NewGoEvaluator()

	testCases := []struct {
		expression string
		errorType  domain.ErrorType
	}{
		{"2 +", domain.ErrorTypeInvalidExpression},
		{"* 5", domain.ErrorTypeInvalidExpression},
		{"(2 + 2", domain.ErrorTypeInvalidExpression},
		{"2 + 2)", domain.ErrorTypeInvalidExpression},
	}

	for _, tc := range testCases {
		// Act
		_, err := evaluator.Evaluate(domain.Expression(tc.expression))

		// Assert
		if err == nil {
			t.Errorf("Expected error for invalid expression '%s', got nil", tc.expression)
			continue
		}

		if err.Type != tc.errorType {
			t.Errorf("For expression '%s', expected error type to be %s, got %s", tc.expression, tc.errorType, err.Type)
		}
	}
}

func TestEvaluate_DivisionByZero(t *testing.T) {
	// Arrange
	evaluator := evaluator.NewGoEvaluator()
	expression := "10 / 0"

	// Act
	ans, err := evaluator.Evaluate(domain.Expression(expression))
	log.Println("Result:", ans)
	// Assert
	if err == nil {
		t.Error("Expected error for division by zero, got nil")
		return
	}

	if err.Type != domain.ErrorTypeDivisionByZero {
		t.Errorf("Expected error type to be %s, got %s", domain.ErrorTypeDivisionByZero, err.Type)
	}
}
