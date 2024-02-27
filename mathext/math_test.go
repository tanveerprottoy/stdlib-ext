package mathext

import "testing"

func TestPercentage(t *testing.T) {
	tests := []struct {
		name     string
		num      int64
		denom    int64
		expected float64
	}{
		{"50 percent", 50, 100, 50.0},
		{"25 percent", 1, 4, 25.0},
		{"0 percent", 0, 100, 0.0},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Percentage(tt.num, tt.denom)
			if actual != tt.expected {
				t.Errorf("Percentage(%d, %d) = %v; want %v", tt.num, tt.denom, actual, tt.expected)
			}
		})
	}
}
