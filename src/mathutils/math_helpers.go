// Package mathutils provides some mathematical utilities used in the raytracer.
package mathutils

import "math"

// Return the radians corresponding to the given angle(in degrees).
func ToRadians(angle float64) float64 {
	return angle / 180.0 * math.Pi
}

// Return the degrees corresponding to the given angle(in radians).
func ToDegrees(angle float64) float64 {
	return angle * 180.0 / math.Pi
}
