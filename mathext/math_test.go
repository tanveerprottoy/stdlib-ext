package mathext

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		val0 int
		val1 int
		exp  int
	}{
		{"2 + 5", 2, 5, 7},
		{"9 + 5", 9, 5, 14},
		{"27 + 45", 27, 45, 72},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Add(tt.val0, tt.val1)
			if actual != tt.exp {
				t.Errorf("Add(%d, %d) = %v; want %v", tt.val0, tt.val1, actual, tt.exp)
			}
		})
	}
}

func TestPercentage(t *testing.T) {
	tests := []struct {
		name  string
		num   int64
		denom int64
		exp   float64
	}{
		{"50 percent", 50, 100, 50.0},
		{"25 percent", 1, 4, 25.0},
		{"0 percent", 0, 100, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Percentage(tt.num, tt.denom)
			if actual != tt.exp {
				t.Errorf("Percentage(%d, %d) = %v; want %v", tt.num, tt.denom, actual, tt.exp)
			}
		})
	}
}
