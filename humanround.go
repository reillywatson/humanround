// package humanround applies heuristic rules to round numbers to a more "human" value.
package humanround

import (
	"math"
	"strconv"
	"strings"
)

type Unit string

const Inch Unit = "inch"

type opts struct {
	unit Unit
}

type Option func(o *opts)

func WithUnit(unit Unit) Option {
	return func(o *opts) { o.unit = unit }
}

// Round rounds f to a more human-readable number. Its heuristics depend on the size of the number,
// and the provided options. It's intended for "human-scale" numbers, ie numbers that might plausibly be used to describe a human.
func Round(f float64, options ...Option) float64 {
	o := &opts{}
	for _, opt := range options {
		opt(o)
	}
	var precision int
	switch {
	case f < 1:
		precision = 3
	case f < 10:
		precision = 2
	case f < 100:
		precision = 1
	default:
		precision = int(2 - math.Log10(f))
	}
	// prefer "round" numbers if they're close enough
	if precision < 2 && strings.HasSuffix(strconv.Itoa(int(math.Floor(f))), "0") {
		f = math.Floor(f)
	}
	if precision < 2 && strings.HasSuffix(strconv.Itoa(int(math.Ceil(f))), "0") {
		f = math.Ceil(f)
	}
	switch o.unit {
	case Inch:
		f = roundInches(f, precision)
	}
	return roundToPrecision(f, precision)
}

// inches are a non-decimal measurement: people like to think about them in halves, fourths, and eighths
func roundInches(f float64, precision int) float64 {
	floor := math.Floor(f)
	if precision == 1 {
		// round to half
		return nearest(f, floor, floor+.5, floor+1)
	}
	if precision == 0 {
		return nearest(f, floor, floor+1)
	}
	if precision == 2 {
		return nearest(f, floor, floor+.25, floor+.5, floor+.75, floor+1)
	}
	return nearest(f, floor, floor+.125, floor+.25, floor+.375, floor+.5, floor+.625, floor+.75, floor+.875, floor+1)
}

func nearest(val float64, targets ...float64) float64 {
	if len(targets) == 0 {
		return val
	}
	res := targets[0]
	diff := math.MaxFloat64
	for _, t := range targets {
		if d := math.Abs(val - t); d < diff {
			res = t
			diff = d
		}
	}
	return res
}

func roundToPrecision(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
