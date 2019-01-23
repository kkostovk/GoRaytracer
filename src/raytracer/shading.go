package raytracer

import (
	"GoRaytracer/src/mathutils"
	"GoRaytracer/src/utils"
	"math"
)

type Texture interface {
	Sample(info *IntersectionInfo) utils.Color
}

type Shader interface {
	Shade(*Ray, *IntersectionInfo) utils.Color
}

type SimpleColor struct {
	color utils.Color
}

func NewSimpleColor(color utils.Color) SimpleColor {
	return SimpleColor{color}
}

func (s *SimpleColor) Sample(_info *IntersectionInfo) utils.Color {
	return s.color
}

type Checker struct {
	color1 utils.Color
	color2 utils.Color
	scale  float64
}

func NewChecker(color1, color2 utils.Color, scale float64) Checker {
	return Checker{color1, color2, scale}
}

func (c *Checker) Sample(info *IntersectionInfo) utils.Color {
	x := int(math.Floor(info.U * c.scale / 5.0))
	y := int(math.Floor(info.V * c.scale / 5.0))
	if (x+y)%2 == 0 {
		return c.color1
	} else {
		return c.color2
	}
}

type Lambert struct {
	color   utils.Color
	texture *Texture
}

func NewLambert(color utils.Color, texture *Texture) Lambert {
	return Lambert{color, texture}
}

func (l *Lambert) Shade(ray *Ray, info *IntersectionInfo) utils.Color {
	diffuse := l.color
	if l.texture != nil {
		diffuse = (*l.texture).Sample(info)
	}

	v1 := info.Normal
	v2 := mathutils.VectorSubstraction(Light1.position, info.Position)
	distanceToLightSqr := v2.LengthSqr()
	v2.Normalize()

	lambertCoeff := mathutils.DotProduct(v1, v2)
	attenuationCoeff := 1.0 / distanceToLightSqr
	diffuse.Multiply(lambertCoeff * attenuationCoeff * Light1.power)
	return diffuse
}
