package mathutils

import "math"

func ToRadians(angle float64) float64 {
	return angle / 180.0 * math.Pi
}

func ToDegrees(angle float64) float64 {
	return angle * 180.0 / math.Pi
}
