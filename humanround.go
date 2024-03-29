// package humanround applies heuristic rules to round numbers to a more "human" value.
package humanround

import (
	"math"
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
	precision := 3 - int(math.Ceil(math.Log10(f)))
	// prefer "round" numbers if they're close enough
	if roundUpToPrecision(f, precision-2) == roundUpToPrecision(f, precision-1) {
		f = roundUpToPrecision(f, precision-1)
	} else if roundDownToPrecision(f, precision-2) == roundDownToPrecision(f, precision-1) {
		f = roundDownToPrecision(f, precision-1)
	}
	switch o.unit {
	case Inch:
		f = roundInches(f, precision)
	}
	return roundNearestToPrecision(f, precision)
}

// inches are a non-decimal measurement: people like to think about them in halves, fourths, and eighths
func roundInches(f float64, precision int) float64 {
	floor := math.Floor(f)
	switch precision {
	case 1:
		// round to half
		return nearest(f, floor, floor+.5, floor+1)
	case 2:
		// round to fourths
		return nearest(f, floor, floor+.25, floor+.5, floor+.75, floor+1)
	case 3:
		// round to eighths
		return nearest(f, floor, floor+.125, floor+.25, floor+.375, floor+.5, floor+.625, floor+.75, floor+.875, floor+1)
	default:
		return f
	}
}

func nearest(val float64, targets ...float64) float64 {
	res := val
	diff := math.MaxFloat64
	for _, t := range targets {
		if d := math.Abs(val - t); d < diff {
			res = t
			diff = d
		}
	}
	return res
}

func roundNearestToPrecision(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}

func roundDownToPrecision(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift) / shift
}
func roundUpToPrecision(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Ceil(f*shift) / shift
}
