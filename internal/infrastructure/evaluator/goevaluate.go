package evaluator

import (
	"arithmetic-telegram-bot/internal/domain"
	"github.com/Knetic/govaluate"
	"log"
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
	if expression == nil {
		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeInvalidExpression,
			Message: "Expression is nil",
			Err:     nil,
		}
	}

	result, err := expression.Evaluate(nil)
	log.Println("result:", result)
	if err != nil {
		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeUnknown,
			Message: "Unknown error during expression evaluation",
			Err:     err,
		}
	}

	if result == nil {
		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeInvalidExpression,
			Message: "Expression evaluation returned nil",
			Err:     nil,
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
