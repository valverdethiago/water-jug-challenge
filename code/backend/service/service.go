package service

import "fmt"

const (
	fillBucketPattern    = "Fill bucket %s"
	emptyBucketPattern   = "Empty bucket %s"
	transferBucketFormat = "Transfer bucket %s to bucket %s"
)

type WaterJugService interface {
	SolveWaterJugProblem(xCap, yCap, target int) ([]State, bool, error)
}

type WaterJugServiceImpl struct {
}

func NewWaterJugServiceImpl() WaterJugService {
	return &WaterJugServiceImpl{}
}

func (s *WaterJugServiceImpl) SolveWaterJugProblem(xCap, yCap, target int) ([]State, bool, error) {
	if err := s.validateInput(xCap, yCap, target); err != nil {
		return nil, false, err
	}
	var result []State

	currentState := newState(0, 0, "Start")
	stepCount := 0

	for !isGoalState(currentState, target) {
		stepCount++
		if currentState.BucketX == 0 {
			currentState.BucketX = xCap
			currentState.Explanation = fmt.Sprintf(fillBucketPattern, "X")
		} else if currentState.BucketY == yCap {
			currentState.BucketY = 0
			currentState.Explanation = fmt.Sprintf(emptyBucketPattern, "Y")
		} else {
			transfer := pour(currentState.BucketX, currentState.BucketY, yCap)
			currentState.BucketX -= transfer
			currentState.BucketY += transfer
			currentState.Explanation = fmt.Sprintf(transferBucketFormat, "X", "Y")
		}
		result = append(result, currentState)
	}
	return result, true, nil
}

func (s *WaterJugServiceImpl) validateInput(xCap int, yCap int, target int) error {
	if target <= 0 || yCap <= 0 || xCap <= 0 {
		return &ValidationError{NegativeNumbers}
	}
	if target > xCap && target > yCap {
		return &ValidationError{ZGreaterThanXAndY}
	}
	if target%gcd(xCap, yCap) != 0 {
		return &ValidationError{TargetNotDivisibleByGcd}
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// get gcd - the greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to check if the current state is the goal state
func isGoalState(state State, target int) bool {
	return state.BucketX == target || state.BucketY == target
}

// Function to simulate the pouring of water from one jug to another
func pour(from, to, sizeTo int) int {
	transfer := min(from, sizeTo-to)
	return transfer
}
