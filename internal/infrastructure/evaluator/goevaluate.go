package evaluator

import (
	"arithmetic-telegram-bot/internal/domain"
	"github.com/Knetic/govaluate"
	"math"
)

type goEvaluator struct{}

func NewGoEvaluator() domain.ExpressionEvaluator {
	return &goEvaluator{}
}

func (e *goEvaluator) Evaluate(expr domain.Expression) (domain.Answer, *domain.CalculatorError) {
	expression, err := govaluate.NewEvaluableExpression(string(expr))
	if err != nil {

		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeInvalidExpression,
			Message: "Invalid expression syntax",
			Err:     err,
		}
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeUnknown,
			Message: "Unknown error during expression evaluation",
			Err:     err,
		}
	}

	if floatResult, ok := result.(float64); ok {
		if math.IsInf(floatResult, 0) {
			return 0, &domain.CalculatorError{
				Type:    domain.ErrorTypeDivisionByZero,
				Message: "Division by zero error",
				Err:     nil,
			}
		}
		return domain.Answer(floatResult), nil
	}

	return 0, &domain.CalculatorError{
		Type:    domain.ErrorTypeUnknown,
		Message: "Unexpected result type",
		Err:     nil,
	}
}
