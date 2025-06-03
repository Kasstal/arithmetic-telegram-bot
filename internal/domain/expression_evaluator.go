package domain

type ExpressionEvaluator interface {
	Evaluate(expr Expression) (Answer, *CalculatorError)
}
