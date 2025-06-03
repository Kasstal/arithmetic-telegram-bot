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
