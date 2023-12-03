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

func GenerateSuccessors(s State, xCap, yCap int) []State {
	var successors []State

	// Fill X
	if s.BucketX < xCap {
		successors = append(successors, newState(xCap, s.BucketY, fmt.Sprintf(fillBucketPattern, "X"), &s))
	}

	// Fill Y
	if s.BucketY < yCap {
		successors = append(successors, newState(s.BucketX, yCap, fmt.Sprintf(fillBucketPattern, "Y"), &s))
	}

	// Empty X
	if s.BucketX > 0 {
		successors = append(successors, newState(0, s.BucketY, fmt.Sprintf(emptyBucketPattern, "X"), &s))
	}

	// Empty Y
	if s.BucketY > 0 {
		successors = append(successors, newState(s.BucketY, 0, fmt.Sprintf(emptyBucketPattern, "Y"), &s))
	}

	// Transfer X to Y
	transfer := min(s.BucketX, yCap-s.BucketY)
	if transfer > 0 {
		successors = append(successors, newState(s.BucketX-transfer, s.BucketY+transfer, fmt.Sprintf(transferBucketFormat, "X", "Y"), &s))
	}

	// Transfer Y to X
	transfer = min(s.BucketY, xCap-s.BucketX)
	if transfer > 0 {
		successors = append(successors, newState(s.BucketX+transfer, s.BucketY-transfer, fmt.Sprintf(transferBucketFormat, "Y", "X"), &s))
	}

	return successors
}

func (s *WaterJugServiceImpl) SolveWaterJugProblem(xCap, yCap, target int) ([]State, bool, error) {
	if err := s.validateInput(xCap, yCap, target); err != nil {
		return nil, false, err
	}

	start := newState(0, 0, "Start", nil)
	queue := []State{start}
	visited := make(map[State]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check if we have reached the target in either jug
		if current.BucketX == target || current.BucketY == target {
			// Check if the other jug is empty and add an extra step if needed
			if current.BucketX == target && current.BucketY != 0 {
				extraStep := newState(target, 0, fmt.Sprintf(emptyBucketPattern, "Y"), &current)
				return append(tracePath(current), extraStep), true, nil
			} else if current.BucketY == target && current.BucketX != 0 {
				extraStep := newState(0, target, fmt.Sprintf(emptyBucketPattern, "X"), &current)
				return append(tracePath(current), extraStep), true, nil
			}
			return tracePath(current), true, nil
		}

		for _, next := range GenerateSuccessors(current, xCap, yCap) {
			if !visited[next] && next.isValidState(xCap, yCap) {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	return nil, false, nil
}

func tracePath(s State) []State {
	var path []State
	for s.Previous != nil {
		path = append([]State{s}, path...)
		s = *s.Previous
	}
	return path
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

// get gcd - greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
