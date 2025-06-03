package tests

import (
	"arithmetic-telegram-bot/internal/app/usecase"
	"arithmetic-telegram-bot/internal/domain"
	"testing"
)

// MockExpressionEvaluator мок для интерфейса ExpressionEvaluator
type MockExpressionEvaluator struct {
	EvaluateFunc func(expr domain.Expression) (domain.Answer, *domain.CalculatorError)
}

func (m *MockExpressionEvaluator) Evaluate(expr domain.Expression) (domain.Answer, *domain.CalculatorError) {
	return m.EvaluateFunc(expr)
}

func TestCalculate_ValidExpression(t *testing.T) {
	// Arrange
	mockEvaluator := &MockExpressionEvaluator{
		EvaluateFunc: func(expr domain.Expression) (domain.Answer, *domain.CalculatorError) {
			return domain.Answer(10), nil
		},
	}

	calculatorService := usecase.NewCalculatorService(mockEvaluator)

	// Act
	result, err := calculatorService.Calculate("2 + 8")

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result != 10.0 {
		t.Errorf("Expected result to be 10.0, got %f", result)
	}
}

func TestCalculate_EmptyExpression(t *testing.T) {
	// Arrange
	mockEvaluator := &MockExpressionEvaluator{}
	calculatorService := usecase.NewCalculatorService(mockEvaluator)

	// Act
	_, err := calculatorService.Calculate("")

	// Assert
	if err == nil {
		t.Error("Expected error for empty expression, got nil")
	}

	if err.Type != domain.ErrorTypeInvalidExpression {
		t.Errorf("Expected error type to be %s, got %s", domain.ErrorTypeInvalidExpression, err.Type)
	}
}

func TestCalculate_EvaluationError(t *testing.T) {
	// Arrange
	mockEvaluator := &MockExpressionEvaluator{
		EvaluateFunc: func(expr domain.Expression) (domain.Answer, *domain.CalculatorError) {
			return 0, &domain.CalculatorError{
				Type:    domain.ErrorTypeDivisionByZero,
				Message: "Division by zero error",
			}
		},
	}

	calculatorService := usecase.NewCalculatorService(mockEvaluator)

	// Act
	_, err := calculatorService.Calculate("1/0")

	// Assert
	if err == nil {
		t.Error("Expected error for division by zero, got nil")
	}

	if err.Type != domain.ErrorTypeDivisionByZero {
		t.Errorf("Expected error type to be %s, got %s", domain.ErrorTypeDivisionByZero, err.Type)
	}
}
