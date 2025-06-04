package usecase

import (
	"arithmetic-telegram-bot/internal/domain"
	"strings"
)

type CalculatorService interface {
	Calculate(expr string) (float64, *domain.CalculatorError)
}

type calculatorService struct {
	evaluator domain.ExpressionEvaluator
}

func (c *calculatorService) Calculate(expr string) (float64, *domain.CalculatorError) {
	cleanedExpr := strings.TrimSpace(expr)
	if cleanedExpr == "" {
		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeInvalidExpression,
			Message: "Expression cannot be empty",
		}
	}

	if !c.isCorrectExpression(cleanedExpr) {
		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeInvalidExpression,
			Message: "Invalid expression syntax",
		}
	}

	result, err := c.evaluator.Evaluate(domain.Expression(cleanedExpr))
	if err != nil {
		return 0, err
	}
	return float64(result), nil
}

func NewCalculatorService(evaluator domain.ExpressionEvaluator) CalculatorService {
	return &calculatorService{
		evaluator: evaluator,
	}
}

func (c *calculatorService) isCorrectExpression(expr string) bool {

	if strings.TrimSpace(expr) == "" {
		return false
	}

	// Check for invalid characters (only allow digits, operators, and parentheses)
	for _, char := range expr {
		if !isValidCharacter(char) {
			return false
		}
	}

	return true
}

func isValidCharacter(char rune) bool {
	if (char >= '0' && char <= '9') ||
		char == '+' || char == '-' ||
		char == '*' || char == '/' ||
		char == '%' ||
		char == '(' || char == ')' ||
		char == '&' || char == '|' ||
		char == '^' ||
		char == '<' || char == '>' ||
		char == ' ' {
		return true
	}
	return false
}
