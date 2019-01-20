package raytracer

import (
	"GoRaytracer/src/utils"
	"math"
)

type Shader interface {
	Shade(*Ray, *IntersectionInfo) utils.Color
}

type Checker struct {
	color1 utils.Color
	color2 utils.Color
}

func NewChecker(color1, color2 utils.Color) Checker {
	return Checker{color1, color2}
}

func (c *Checker) Shade(ray *Ray, info *IntersectionInfo) utils.Color {
	x := int(math.Floor(info.U / 5.0))
	y := int(math.Floor(info.V / 5.0))

	if (x+y)%2 == 0 {
		return c.color1
	} else {
		return c.color2
	}
}
