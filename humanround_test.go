package humanround

import "testing"

func TestRound(t *testing.T) {
	tests := []struct {
		in   float64
		unit Unit
		exp  float64
	}{
		{
			in:  50.8,
			exp: 50,
		},
		{
			in:   2.26796,
			unit: Inch,
			exp:  2.25,
		},
		{
			in:  453.592,
			exp: 454,
		},
	}
	for _, test := range tests {
		if test.unit == "" {
			test.unit = Unknown
		}
		got := Round(test.in, test.unit)
		if got != test.exp {
			t.Errorf("got %f, expected %f", got, test.exp)
		}
	}
}
