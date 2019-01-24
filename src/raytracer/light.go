package raytracer

import (
	"GoRaytracer/src/mathutils"
	"GoRaytracer/src/utils"
)

type Light struct {
	position mathutils.Vector
	color    utils.Color
	power    float64
}

func NewLight(position mathutils.Vector, color utils.Color, power float64) Light {
	return Light{position, color, power}
}

var Light1 Light
var AmbientLight Light
