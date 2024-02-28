package humanround

import "testing"

func TestRound(t *testing.T) {
	tests := []struct {
		in   float64
		opts []Option
		exp  float64
	}{
		{
			in:  50.8,
			exp: 50,
		},
		{
			in:   2.26796,
			opts: []Option{WithUnit(Inch)},
			exp:  2.25,
		},
		{
			in:   0.13,
			opts: []Option{WithUnit(Inch)},
			exp:  0.125,
		},
		{
			in:   32.26796,
			opts: []Option{WithUnit(Inch)},
			exp:  32.5,
		},
		{
			in:  2.26796,
			exp: 2.27,
		},
		{
			in:  453.592,
			exp: 454,
		},
		{
			in:  299,
			exp: 300,
		},
		{
			in:  453.592,
			exp: 454,
		},
		{
			in:  1022,
			exp: 1000,
		},
		{
			in:  5375,
			exp: 5380,
		},
		{
			in:  53750,
			exp: 53800,
		},
		{
			in:  55555,
			exp: 55600,
		},
	}
	for _, test := range tests {
		got := Round(test.in, test.opts...)
		if got != test.exp {
			t.Errorf("got %f, expected %f", got, test.exp)
		}
	}
}

func TestRoundDownToPrecision(t *testing.T) {
	tests := []struct {
		in     float64
		places int
		exp    float64
	}{
		{1.27, 1, 1.2},
		{1.27, 2, 1.27},
		{1.27, 0, 1},
		{1.27, -1, 0},
		{123, -1, 120},
	}
	for _, test := range tests {
		got := roundDownToPrecision(test.in, test.places)
		if got != test.exp {
			t.Errorf("got %f, expected %f", got, test.exp)
		}
	}
}

func TestRoundUpToPrecision(t *testing.T) {
	tests := []struct {
		in     float64
		places int
		exp    float64
	}{
		{1.27, 1, 1.3},
		{1.27, 2, 1.27},
		{1.27, 0, 2},
		{1.27, -1, 10},
		{123, -1, 130},
	}
	for _, test := range tests {
		got := roundUpToPrecision(test.in, test.places)
		if got != test.exp {
			t.Errorf("got %f, expected %f", got, test.exp)
		}
	}
}
