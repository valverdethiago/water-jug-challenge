package service

const (
	NegativeNumbers         = "All input variables should be positive integers"
	ZGreaterThanXAndY       = "Z can't be greater than X and Y"
	TargetNotDivisibleByGcd = "Target is not divisible by Greatest Common Divisor of X and Y"
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
