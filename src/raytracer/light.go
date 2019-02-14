// Package raytracer provides the raytracer logic.
package raytracer

import (
	"GoRaytracer/src/mathutils"
	"GoRaytracer/src/utils"
)

// Light defines a point light.
type Light struct {
	position mathutils.Vector
	color    utils.Color
	power    float64
}

// NewLight creates and return a new point light.
func NewLight(position mathutils.Vector, color utils.Color, power float64) Light {
	return Light{position, color, power}
}
