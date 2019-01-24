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

type Phong struct {
	color              utils.Color
	texture            *Texture
	specularMultiplier float64
	specularExponent   float64
}

func NewPhong(color utils.Color, texture *Texture, specularMultiplier, specularExponent float64) Phong {
	return Phong{color, texture, specularMultiplier, specularExponent}
}

func (p *Phong) Shade(ray *Ray, info *IntersectionInfo) utils.Color {
	diffuse := p.color
	if p.texture != nil {
		diffuse = (*p.texture).Sample(info)
	}

	v1 := info.Normal
	v2 := mathutils.VectorSubstraction(Light1.position, info.Position)
	distanceToLightSqr := v2.LengthSqr()
	v2.Normalize()

	lambertCoeff := mathutils.DotProduct(v1, v2)
	fromLight := Light1.power / distanceToLightSqr

	reflected := mathutils.Reflect(mathutils.VectorSubstraction(info.Position, Light1.position), info.Normal)
	toCamera := ray.Direction
	toCamera.UnaryMinus()
	cosGamma := mathutils.DotProduct(toCamera, reflected)
	phongCoeff := 0.0
	if cosGamma > 0 {
		phongCoeff = math.Pow(cosGamma, p.specularExponent)
	}

	diffuse.Multiply(lambertCoeff * fromLight)
	result := utils.NewColor(255, 255, 255)
	result.Multiply(phongCoeff * p.specularMultiplier * fromLight)

	return utils.ColorAddition(result, diffuse)
}
