// Package raytracer provides the raytracer logic.
package raytracer

import (
	"GoRaytracer/src/mathutils"
	"GoRaytracer/src/utils"
	"math"
)

// Texture provides an interface for sampling textures.
type Texture interface {
	Sample(info *IntersectionInfo) utils.Color
}

// Shader provides an interface for shading a surface.
type Shader interface {
	Shade(*Ray, *IntersectionInfo, *Scene) utils.Color
}

// SimpleColor defines a simple color texture.
type SimpleColor struct {
	color utils.Color // The color of the simple color texture.
}

// NewSimpleColor creates and returns a SimpleColor.
func NewSimpleColor(color utils.Color) SimpleColor {
	return SimpleColor{color}
}

// Sample implements sampling for SimpleColor.
func (s *SimpleColor) Sample(_info *IntersectionInfo) utils.Color {
	return s.color
}

// Checker defines the checker texture.
type Checker struct {
	color1 utils.Color
	color2 utils.Color
	scale  float64
}

// NewChecker creates and returns a new checker texture.
func NewChecker(color1, color2 utils.Color, scale float64) Checker {
	return Checker{color1, color2, scale}
}

// Sample implements sampling for Checker.
func (c *Checker) Sample(info *IntersectionInfo) utils.Color {
	x := int(math.Floor(info.U * c.scale / 5.0))
	y := int(math.Floor(info.V * c.scale / 5.0))
	if (x+y)%2 == 0 {
		return c.color1
	}

	return c.color2
}

// Lambert defines a lambert shader.
type Lambert struct {
	color   utils.Color
	texture *Texture
}

// NewLambert creates and returns a new lambert shader.
func NewLambert(color utils.Color, texture Texture) Lambert {
	return Lambert{color, &texture}
}

// Shade implements a lambert shader.
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

// Phong defines a phong shader.
type Phong struct {
	color              utils.Color
	texture            *Texture
	specularMultiplier float64
	specularExponent   float64
}

// NewPhong creates and returns a new phong shader.
func NewPhong(color utils.Color, texture Texture, specularMultiplier, specularExponent float64) Phong {
	return Phong{color, &texture, specularMultiplier, specularExponent}
}

// Shade implements a phong shader.
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

// visibilityCheck checks if there is an object between start and end.
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

// getLightContribution return the light contribution for the current point.
func getLightContribution(ray *Ray, info *IntersectionInfo, light *Light) utils.Color {
	vectorToLight := mathutils.VectorSubstraction(light.position, info.Position)
	distanceToLightSqr := vectorToLight.LengthSqr()

	vectorToLight.Normalize()
	cosTheta := mathutils.DotProduct(vectorToLight, mathutils.Faceforward(ray.Direction, info.Normal))

	result := light.color
	result.Multiply(light.power / distanceToLightSqr * cosTheta)

	return result
}
