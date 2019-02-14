// Package raytracer provides the raytracer logic.
package raytracer

import (
	"GoRaytracer/src/mathutils"
)

// Ray defines a ray in the 3-dimentional space.
type Ray struct {
	Start, Direction mathutils.Vector
}

// NewRay creates and returns a new ray.
func NewRay(start, direction mathutils.Vector) Ray {
	return Ray{start, direction}
}
