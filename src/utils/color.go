// Package utils provides some simple utilities for the raytracer.
package utils

// Defines a color.
type Color [3]float64

// Crate a new color with the given red, green and blue components and return it.
func NewColor(r, g, b uint8) Color {
	return Color{float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0}
}

// Return the red, green and blue components of c.
func (c *Color) ToRGB() (uint8, uint8, uint8) {
	clamp := func(val float64) float64 {
		if val < 0 {
			return 0
		}
		if val > 1 {
			return 1
		}
		return val
	}

	return uint8(clamp(c[0]) * 255), uint8(clamp(c[1]) * 255), uint8(clamp(c[2]) * 255)
}

// Multiply the components of c by multiplier.
func (c *Color) Multiply(multiplier float64) {
	*c = MultiplyColorFloat(*c, multiplier)
}

// Divide the components of c by divider.
func (c *Color) Divide(divider float64) {
	*c = DivideColorFloat(*c, divider)
}

// Add lhs and rhs and return the color with the result.
func ColorAddition(lhs, rhs Color) Color {
	return Color{lhs[0] + rhs[0], lhs[1] + rhs[1], lhs[2] + rhs[2]}
}

// Multiply lhs and rhs and return the color with the result.
func ColorMultiplication(lhs, rhs Color) Color {
	return Color{lhs[0] * rhs[0], lhs[1] * rhs[1], lhs[2] * rhs[2]}
}

// Multiply the components of c by multiplier and return the color with the result.
func MultiplyColorFloat(c Color, multiplier float64) Color {
	return Color{c[0] * multiplier, c[1] * multiplier, c[2] * multiplier}
}

// Divide the components of c by divider and return the color with the result.
func DivideColorFloat(c Color, divider float64) Color {
	return Color{c[0] / divider, c[1] / divider, c[2] / divider}
}
