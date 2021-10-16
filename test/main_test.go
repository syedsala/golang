package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestCalculateMultiple(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"+ve Number", args {100}, 102},
		{"-ve Number", args {-52}, -50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Calculate(tt.args.x); gotResult != tt.wantResult {
				t.Errorf("Calculate() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}