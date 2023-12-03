package service

import (
	"reflect"
	"testing"
)

type args struct {
	x int
	y int
	z int
}

func TestNewWaterJugServiceImpl(t *testing.T) {
}

func TestWaterJugServiceImpl_SolveWaterJugProblem(t *testing.T) {
	tests := []struct {
		name      string
		args      args
		want      []State
		wantBool  bool
		wantError error
	}{
		{
			name:      "Should fail with negative numbers",
			args:      args{-1, -1, -1},
			wantBool:  false,
			wantError: &ValidationError{NegativeNumbers},
		},
		{
			name:      "Should fail with Z greater than X and Y",
			args:      args{45, 35, 50},
			wantBool:  false,
			wantError: &ValidationError{ZGreaterThanXAndY},
		},
		{
			name: "Should work with the basic test scenario",
			args: args{2, 10, 4},
			want: []State{
				newState(2, 0, "Fill bucket X", nil),
				newState(0, 2, "Transfer bucket X to bucket Y", nil),
				newState(2, 2, "Fill bucket X", nil),
				newState(0, 4, "Transfer bucket X to bucket Y", nil),
			},
			wantBool:  true,
			wantError: nil,
		},
		{
			name: "Should empty latest bucket after finding a solution",
			args: args{4, 3, 2},
			want: []State{
				newState(0, 3, "Fill bucket Y", nil),
				newState(3, 0, "Empty bucket Y", nil),
				newState(3, 3, "Fill bucket Y", nil),
				newState(4, 2, "Transfer bucket Y to bucket X", nil),
				newState(0, 2, "Empty bucket X", nil),
			},
			wantBool:  true,
			wantError: nil,
		},
		{
			name:      "Validation should fail when Z is not a multiple of the GCD of x and y",
			args:      args{8, 6, 5},
			wantBool:  false,
			wantError: &ValidationError{TargetNotDivisibleByGcd},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			steps, solution, err := NewWaterJugServiceImpl().SolveWaterJugProblem(tt.args.x, tt.args.y, tt.args.z)
			if !reflect.DeepEqual(err, tt.wantError) {
				t.Errorf("%s should have thrown an error %s", tt.name, tt.wantError)
			}
			if solution != tt.wantBool {
				t.Errorf("%s unexpected solution %v", tt.name, tt.wantBool)
			}
			if len(steps) != len(tt.want) {
				t.Errorf("%s unexpected number of steps [%d]. The number of steps should be %d",
					tt.name, len(steps), len(tt.want))
			}
			for index, step := range steps {
				wanted := &tt.want[index]
				if !wanted.isEquals(step) {
					t.Errorf("%s - Step %d is not as expected. Expected [ %v ], got [ %v] ", tt.name,
						index, wanted, step)
				}
			}
		})
	}
}

func TestWaterJugServiceImpl_validateInput(t *testing.T) {
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		expectedMessage string
	}{
		{
			name:            "Should fail with negative numbers",
			args:            args{-1, -1, -1},
			wantErr:         true,
			expectedMessage: NegativeNumbers,
		},
		{
			name:            "Should fail with Z greater than X and Y",
			args:            args{45, 35, 50},
			wantErr:         true,
			expectedMessage: ZGreaterThanXAndY,
		},
		{
			name:            "Validation should fail when Z is not a multiple of the GCD of x and y",
			args:            args{8, 6, 5},
			wantErr:         true,
			expectedMessage: TargetNotDivisibleByGcd,
		},
		{
			name:    "Should work with the basic test scenario",
			args:    args{2, 10, 4},
			wantErr: false,
		},
		{
			name:    "Should empty latest bucket after finding a solution",
			args:    args{4, 3, 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &WaterJugServiceImpl{}
			if err := s.validateInput(tt.args.x, tt.args.y, tt.args.z); (err != nil) != tt.wantErr {
				t.Errorf("validateInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
