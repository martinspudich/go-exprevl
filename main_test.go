package exprevl

import (
	"errors"
	"testing"
)

func TestEvaluate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		expr          string
		args          map[string]float64
		want          float64
		expectedError error
	}{
		{
			name:          "Test Add",
			expr:          "x + y",
			args:          map[string]float64{"x": 1, "y": 2},
			want:          3,
			expectedError: nil,
		},
		{
			name:          "Test Add",
			expr:          "3 + 4",
			args:          map[string]float64{},
			want:          7,
			expectedError: nil,
		},
		{
			name:          "Test Add",
			expr:          "3 + 4 + 5",
			args:          map[string]float64{},
			want:          12,
			expectedError: nil,
		},
		{
			name:          "Test Add error",
			expr:          "c + y",
			args:          map[string]float64{"x": 1, "y": 2},
			want:          0,
			expectedError: ErrArgNotFound,
		},
		{
			name:          "Test Sub",
			expr:          "x - y",
			args:          map[string]float64{"x": 1, "y": 2},
			want:          -1,
			expectedError: nil,
		},
		{
			name:          "Test Sub",
			expr:          "3 - 2",
			args:          map[string]float64{},
			want:          1,
			expectedError: nil,
		},
		{
			name:          "Test Sub",
			expr:          "14 - 4 - 5",
			args:          map[string]float64{},
			want:          5,
			expectedError: nil,
		},
		{
			name:          "Test Sub error",
			expr:          "c - y",
			args:          map[string]float64{"x": 1, "y": 2},
			want:          0,
			expectedError: ErrArgNotFound,
		},
		{
			name:          "Test Multi",
			expr:          "x * y",
			args:          map[string]float64{"x": 1, "y": 2},
			want:          2,
			expectedError: nil,
		},
		{
			name:          "Test Multi",
			expr:          "3 * 2",
			args:          map[string]float64{},
			want:          6,
			expectedError: nil,
		},
		{
			name:          "Test Multi",
			expr:          "1 * 4 * 5",
			args:          map[string]float64{},
			want:          20,
			expectedError: nil,
		},
		{
			name:          "Test Multi error",
			expr:          "c * y",
			args:          map[string]float64{"x": 1, "y": 2},
			want:          0,
			expectedError: ErrArgNotFound,
		},
		{
			name:          "Test Div",
			expr:          "x / y",
			args:          map[string]float64{"x": 1, "y": 2},
			want:          0.5,
			expectedError: nil,
		},
		{
			name:          "Test Div",
			expr:          "3 / 2",
			args:          map[string]float64{},
			want:          1.5,
			expectedError: nil,
		},
		{
			name:          "Test Div",
			expr:          "10 / 2 / 5",
			args:          map[string]float64{},
			want:          1,
			expectedError: nil,
		},
		{
			name:          "Test Div error",
			expr:          "c / y",
			args:          map[string]float64{"x": 1, "y": 2},
			want:          0,
			expectedError: ErrArgNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate(tt.expr, tt.args)
			if err != nil {
				if !errors.Is(err, tt.expectedError) {
					t.Errorf("Expect error: %s, but got %s", tt.expectedError, err)
					return
				}
			}

			if tt.expectedError != nil {
				if err == nil {
					t.Errorf("Expecting error %s, but got nil", tt.expectedError)
					return
				}
			} else {
				if got != tt.want {
					t.Errorf("expression: %s; got: %v; want: %v", tt.expr, got, tt.want)
				}
			}
		})
	}
}
