package raytracer

import (
	"GoRaytracer/src/mathutils"
)

type Ray struct {
	Start, Direction mathutils.Vector
}

func NewRay(start, direction mathutils.Vector) Ray {
	return Ray{start, direction}
}
