// package humanround applies heuristic rules to round numbers to a more "human" value.
package humanround

import (
	"math"
	"strconv"
	"strings"
)

type Unit string

const Inch Unit = "inch"
const Cm Unit = "cm"
const Pound Unit = "pound"
const Kg Unit = "kg"
const Unknown Unit = "unknown"

// Round rounds f to a more human-readable number. Its heuristics depend on the size of the number,
// and the provided unit. It's intended for "human-scale" numbers, ie numbers that might plausibly be used to describe a human.
func Round(f float64, unit Unit) float64 {
	precision := 2 - int(math.Log10(f))
	// prefer "round" numbers if they're close enough
	if precision < 2 && strings.HasSuffix(strconv.Itoa(int(math.Floor(f))), "0") {
		return math.Floor(f)
	}
	if precision < 2 && strings.HasSuffix(strconv.Itoa(int(math.Ceil(f))), "0") {
		return math.Ceil(f)
	}
	if unit == Inch {
		f = roundInches(f, precision)
	}
	return roundToPrecision(f, precision)
}

// inches are a non-decimal measurement, people like to think about them in halves, fourths, and eighths
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
