package utils

type Color [3]float64

func NewColor(r, g, b uint8) Color {
	return Color{float64(r) / 255.0, float64(g) / 255.0, float64(b) / 255.0}
}

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

func (c *Color) Multiply(multiplier float64) {
	*c = MultiplyColorFloat(*c, multiplier)
}

func (c *Color) Divide(divider float64) {
	*c = DivideColorFloat(*c, divider)
}

func ColorAddition(lhs, rhs Color) Color {
	return Color{lhs[0] + rhs[0], lhs[1] + rhs[1], lhs[2] + rhs[2]}
}

func ColorMultiplication(lhs, rhs Color) Color {
	return Color{lhs[0] * rhs[0], lhs[1] * rhs[1], lhs[2] * rhs[2]}
}

func MultiplyColorFloat(c Color, multiplier float64) Color {
	return Color{c[0] * multiplier, c[1] * multiplier, c[2] * multiplier}
}

func DivideColorFloat(c Color, divider float64) Color {
	return Color{c[0] / divider, c[1] / divider, c[2] / divider}
}
