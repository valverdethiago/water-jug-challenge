package service

type State struct {
	BucketX, BucketY int // Current amounts of water in jugs X and Y
	Explanation      string
	Previous         *State
}

func (s *State) isGoalState(Z int) bool {
	return s.BucketX == Z || s.BucketY == Z
}

func (s *State) isValidState(xCap, yCap int) bool {
	return s.BucketX >= 0 && s.BucketX <= xCap && s.BucketY >= 0 && s.BucketY <= yCap
}
func (s *State) isEquals(another State) bool {
	return s.BucketX == another.BucketX && s.BucketY == another.BucketY && s.Explanation == another.Explanation
}

func newState(x, y int, explanation string, previous *State) State {
	return State{
		BucketX:     x,
		BucketY:     y,
		Explanation: explanation,
		Previous:    previous,
	}
}
