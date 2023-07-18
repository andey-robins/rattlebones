package die

import "testing"

func TestDieFromValue(t *testing.T) {
	tests := []struct {
		val int
	}{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
	}

	for _, test := range tests {
		d := NewDieFromValue(test.val)
		if d.Value != test.val {
			t.Errorf("Expected %d, got %d", test.val, d.Value)
		}
	}
}
