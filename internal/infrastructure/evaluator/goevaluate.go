package evaluator

import (
	"arithmetic-telegram-bot/internal/domain"
	"github.com/Knetic/govaluate"
	"strings"
)

type GoEvaluator struct{}

func (e *GoEvaluator) Evaluate(expr domain.Expression) (domain.Answer, error) {
	expression, err := govaluate.NewEvaluableExpression(string(expr))
	if err != nil {
		errmsg := err.Error()
		if strings.Contains(errmsg, "syntax error") {
			return 0, &domain.CalculatorError{
				Type:    domain.ErrorTypeInvalidExpression,
				Message: "Invalid expression syntax",
				Err:     err,
			}
		}
	}
		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeUnknown,
			Message: "Unknown error during expression evaluation",
			Err:     err,
		}
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "division by zero") {
			return 0, &domain.CalculatorError{
				Type:    domain.ErrorTypeDivisionByZero,
				Message: "Division by zero error",
				Err:     err,
			}
		}
		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeUnknown,
			Message: "Unknown error during expression evaluation",
			Err:     err,
		}

		if floatResult, ok := result.(float64); ok {
			return domain.Answer(floatResult), nil
		}

		return 0, &domain.CalculatorError{
			Type:    domain.ErrorTypeUnknown,
			Message: "Unexpected result type",
			Err:     nil,
	}
}

