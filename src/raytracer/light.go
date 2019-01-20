package raytracer

import "GoRaytracer/src/utils"

type Light struct {
	color utils.Color
	power float64
}

func NewLight(color utils.Color, power float64) Light {
	return Light{color, power}
}
