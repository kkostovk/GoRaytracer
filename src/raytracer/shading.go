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
	Shade(*Ray, *IntersectionInfo, *Scene) utils.Color
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

func NewLambert(color utils.Color, texture Texture) Lambert {
	return Lambert{color, &texture}
}

func (l *Lambert) Shade(ray *Ray, info *IntersectionInfo, scene *Scene) utils.Color {
	var result utils.Color
	diffuse := l.color
	if l.texture != nil {
		diffuse = (*l.texture).Sample(info)
	}

	diffuse = utils.ColorMultiplication(diffuse, scene.ambientLight)

	for _, light := range scene.lights {
		displacedRay := mathutils.VectorAddition(info.Position, mathutils.VectorMultiply(info.Normal, 1e-5))
		if visibilityCheck(displacedRay, light.position, scene) {
			lightContribution := getLightContribution(ray, info, &light)
			result = utils.ColorMultiplication(diffuse, lightContribution)
		} else {
			result = utils.ColorMultiplication(diffuse, utils.Color{0.05, 0.05, 0.05})
		}
	}
	return result
}

type Phong struct {
	color              utils.Color
	texture            *Texture
	specularMultiplier float64
	specularExponent   float64
}

func NewPhong(color utils.Color, texture Texture, specularMultiplier, specularExponent float64) Phong {
	return Phong{color, &texture, specularMultiplier, specularExponent}
}

func (p *Phong) Shade(ray *Ray, info *IntersectionInfo, scene *Scene) utils.Color {
	var result utils.Color
	diffuse := p.color
	if p.texture != nil {
		diffuse = (*p.texture).Sample(info)
	}

	for _, light := range scene.lights {
		reflected := mathutils.Reflect(mathutils.VectorSubstraction(info.Position, light.position), info.Normal)
		toCamera := ray.Direction
		toCamera.UnaryMinus()
		cosGamma := mathutils.DotProduct(toCamera, reflected)
		phongCoeff := 0.0
		if cosGamma > 0 {
			phongCoeff = math.Pow(cosGamma, p.specularExponent)
		}

		lightContribution := utils.ColorMultiplication(scene.ambientLight, getLightContribution(ray, info, &light))
		diffuse = utils.ColorAddition(diffuse, lightContribution)
		specular := utils.MultiplyColorFloat(lightContribution, phongCoeff*p.specularMultiplier)
		result = utils.ColorAddition(diffuse, specular)
	}

	result = utils.ColorMultiplication(result, scene.ambientLight)
	return result
}

func visibilityCheck(start, end mathutils.Vector, scene *Scene) bool {
	direction := mathutils.VectorSubstraction(end, start)
	targetDistance := direction.Length()
	direction.Normalize()
	ray := NewRay(start, direction)

	for _, node := range scene.SceneNodes {
		var info IntersectionInfo
		if !(*node.geometry).Intersect(&ray, &info) {
			continue
		}

		if info.Distance < targetDistance {
			return false
		}
	}

	return true
}

func getLightContribution(ray *Ray, info *IntersectionInfo, light *Light) utils.Color {
	vectorToLight := mathutils.VectorSubstraction(light.position, info.Position)
	distanceToLightSqr := vectorToLight.LengthSqr()

	vectorToLight.Normalize()
	cosTheta := mathutils.DotProduct(vectorToLight, mathutils.Faceforward(ray.Direction, info.Normal))

	result := light.color
	result.Multiply(light.power / distanceToLightSqr * cosTheta)

	return result
}
